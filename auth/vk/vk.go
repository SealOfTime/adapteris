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

type vkAuthenticator struct {
	ExtAccStorage domain.ExternalAccountStorage
	UserStorage   domain.UserStorage
	oauth2.Config
}

func NewVkAuthenticator(
	clientId, clientSecret, successRedirectUrl string,
	extAccStorage domain.ExternalAccountStorage,
	userStorage domain.UserStorage,
) *vkAuthenticator {
	return &vkAuthenticator{
		ExtAccStorage: extAccStorage,
		UserStorage:   userStorage,
		Config: oauth2.Config{
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Endpoint:     vk.Endpoint,
			RedirectURL:  successRedirectUrl,
			Scopes:       []string{"email"},
		},
	}
}

func (a *vkAuthenticator) LoginUrl(afterLoginUrl string) string {
	return a.Config.AuthCodeURL(afterLoginUrl)
}
func (a *vkAuthenticator) Authenticate(ctx context.Context, creds OAuth2AccessCodeCredentials) (*domain.User, error) {
	if creds.ProviderName != "Vk" {
		//If not the OAuth2 credentials authenticating or if the requested authentication provider is not the same
		//as this one, skip.
		return nil, nil
	}

	oa2t, err := a.Exchange(ctx, creds.AccessCode)
	if err != nil {
		return nil, ErrExchangeFailed{Cause: err}
	}

	eid, err := a.ExtractExternalId(ctx, *oa2t)
	if err != nil {
		return nil, ErrOA2InvalidToken{Cause: err}
	}

	ea, err := a.ExtAccStorage.FindByExternalId(ctx, eid)
	if err != nil {
		var errNotFound domain.ErrExternalAccountNotFound
		if errors.As(err, &errNotFound) {
			d, err := a.ExtractUserDefaults(ctx, *oa2t)
			if err != nil {
				return nil, err
			}
			return a.UserStorage.UpsertByEmail(ctx, *d)
		}

		return nil, err
	}

	return ea.User, nil
}

func (a *vkAuthenticator) ExtractExternalId(ctx context.Context, t oauth2.Token) (string, error) {
	exIdRaw := t.Extra("user_id")
	if exIdRaw == nil {
		return "", ErrOA2ExtraMissing{Extra: "user_id"}
	}

	exIdInt := int64(exIdRaw.(float64))
	exId := strconv.FormatInt(exIdInt, 10)
	return exId, nil
}

func (a *vkAuthenticator) ExtractUserDefaults(ctx context.Context, t oauth2.Token) (*domain.User, error) {
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
