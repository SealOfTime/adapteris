package vk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/sealoftime/adapteris/domain/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

const (
	apiVersion = "5.131"
)

var (
	getUserProfileUrl, _ = url.Parse("https://api.vk.com/method/users.get")
)

//Service provides integration with VK.com
type Service struct {
	config oauth2.Config
	auth   *user.AuthService
}

func New(
	authService *user.AuthService,
	clientId, clientSecret, successRedirectUrl string,
) *Service {
	return &Service{
		auth: authService,
		config: oauth2.Config{
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Endpoint:     vk.Endpoint,
			RedirectURL:  successRedirectUrl,
			Scopes:       []string{"email"},
		},
	}
}

//GetUserProfile returns detailed information about the user related to the provided accessToken.
func (s *Service) GetUserProfile(
	ctx context.Context,
	accessToken string,
	fields []string,
) (User, error) {
	reqUrl := *getUserProfileUrl
	q := url.Values{
		"v":            []string{apiVersion},
		"access_token": []string{accessToken},
		"fields":       fields,
	}
	reqUrl.RawQuery = q.Encode()

	resp, err := http.Get(reqUrl.String())
	if err != nil {
		return User{}, ErrVkFetchUserDetails{Cause: err}
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("error closing body after requesting user details from vk.com: %+v", err)
		}
	}()

	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, ErrVkFetchUserDetails{Cause: err}
	}

	var payload struct {
		Response []User
	}
	if err = json.Unmarshal(rawData, &payload); err != nil {
		return User{}, ErrVkFetchUserDetails{Cause: err}
	}
	return payload.Response[0], nil
}

type User struct {
	Domain    string `json:"domain"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
