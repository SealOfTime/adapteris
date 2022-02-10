package vk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	domain "github.com/sealoftime/adapteris"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

type AccessCodeCreds struct {
	AccessCode string
}

type Authenticator struct {
	extAccStorage domain.ExternalAccountStorage
	userStorage   domain.UserStorage
	oauth2.Config
}

func NewAuthenticator(
	clientId, clientSecret, successRedirectUrl string,
	extAccStorage domain.ExternalAccountStorage,
	userStorage domain.UserStorage,
) *Authenticator {
	return &Authenticator{
		extAccStorage: extAccStorage,
		userStorage:   userStorage,
		Config: oauth2.Config{
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Endpoint:     vk.Endpoint,
			RedirectURL:  successRedirectUrl,
			Scopes:       []string{"email"},
		},
	}
}

func (a *Authenticator) LoginUrl(afterLoginUrl string) string {
	return a.Config.AuthCodeURL(afterLoginUrl)
}

func (a *Authenticator) Authenticate(ctx context.Context, rawCreds interface{}) (*domain.User, error) {
	creds, ok := rawCreds.(AccessCodeCreds)
	if !ok {
		return nil, nil
	}

	oa2t, err := a.Exchange(ctx, creds.AccessCode)
	if err != nil {
		return nil, ErrExchangeFailed{Cause: err}
	}

	eid, err := a.extractExternalId(ctx, *oa2t)
	if err != nil {
		return nil, ErrOA2InvalidToken{Cause: err}
	}

	ea, err := a.extAccStorage.FindByExternalId(ctx, eid)
	if err != nil {
		var errNotFound domain.ErrExternalAccountNotFound
		if errors.As(err, &errNotFound) {
			d, err := a.extractUserDefaults(ctx, *oa2t)
			if err != nil {
				return nil, err
			}
			return a.userStorage.UpsertByEmail(ctx, *d)
		}

		return nil, err
	}

	return ea.User, nil
}

func (a *Authenticator) extractExternalId(ctx context.Context, t oauth2.Token) (string, error) {
	exIdRaw := t.Extra("user_id")
	if exIdRaw == nil {
		return "", ErrOA2ExtraMissing{Extra: "user_id"}
	}

	exIdInt := int64(exIdRaw.(float64))
	exId := strconv.FormatInt(exIdInt, 10)
	return exId, nil
}

func (a *Authenticator) extractUserDefaults(ctx context.Context, t oauth2.Token) (*domain.User, error) {
	vk, err := fetchUserDetailsForToken(t)
	if err != nil {
		return nil, err
	}

	fullName := fmt.Sprintf("%s %s", vk.FirstName, vk.LastName)
	return &domain.User{
		FullName:  &fullName,
		ShortName: vk.FirstName,
		Email:     t.Extra("email").(string),
		Vk:        &vk.Domain,
	}, nil
}

type vkUserResponse struct {
	Response []vkUser
}

type vkUser struct {
	Domain    string `json:"domain"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func fetchUserDetailsForToken(t oauth2.Token) (*vkUser, error) {
	req, _ := url.Parse("https://api.vk.com/method/users.get")
	q := url.Values{
		"v":            []string{"5.131"},
		"access_token": []string{t.AccessToken},
		"fields":       []string{"domain"},
	}
	req.RawQuery = q.Encode()

	resp, err := http.Get(req.String())
	if err != nil {
		return nil, ErrVkFetchUserDetails{Cause: err}
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrVkFetchUserDetails{Cause: err}
	}

	var payload vkUserResponse
	if err = json.Unmarshal(rawData, &payload); err != nil {
		return nil, ErrVkFetchUserDetails{Cause: err}
	}
	return &payload.Response[0], nil
}
