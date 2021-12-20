package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"log"
	"net/url"
	"strconv"
)

type Config struct {
	CallbackUrl string
	Vk          ProviderConfig
}

type JwtConfig struct {
	Alg           string `mapstructure:"alg"`
	signingMethod jwt.SigningMethod
	Key           string
	key           []byte
}

type ProviderConfig struct {
	Jwt    JwtConfig
	OAuth2 oauth2.Config
}

func setupProviders(cfg *Config, storage ExternalAccountStorage) []Authenticator {
	providers := make([]Authenticator, 0, 3)
	providers = append(providers, setupVK(cfg, storage))
	return providers
}

func setupProviderConfig(cfg *ProviderConfig, defaultCallback string) {
	redirectUrl := cfg.OAuth2.RedirectURL
	if redirectUrl == "" {
		redirectUrl = defaultCallback
	}

	if _, err := url.Parse(redirectUrl); err != nil {
		log.Fatalf("Couldn't initialize authentication provider with redirectUrl '%s': %+v", redirectUrl, err)
	}

	cfg.Jwt.signingMethod = jwt.GetSigningMethod(cfg.Jwt.Alg)
	if cfg.Jwt.signingMethod == nil {
		log.Fatalf("Couldn't initialize authentication provider with algorithm '%s'", cfg.Jwt.Alg)
	}
	cfg.Jwt.key = []byte(cfg.Jwt.Key)
}

func setupVK(cfg *Config, storage ExternalAccountStorage) Authenticator {
	defaultHandler := fmt.Sprintf("%s/vk", cfg.CallbackUrl)
	setupProviderConfig(&cfg.Vk, defaultHandler)

	return NewOAuth2Authenticator("Vk", &cfg.Vk, extIdFromOA2TokenExtra("user_id"), storage)
}

func extIdFromOA2TokenExtra(key string) func(ctx context.Context, oa2t *oauth2.Token) (string, error) {
	return func(ctx context.Context, oa2t *oauth2.Token) (string, error) {
		exIdRaw := oa2t.Extra(key)
		if exIdRaw == nil {
			return "", ErrOA2ExtraMissing{Extra: key}
		}
		var exId string
		switch exIdRaw.(type) {
		case string:
			exId = exIdRaw.(string)
		case int:
			{
				exIdInt := exIdRaw.(int)
				exId = strconv.Itoa(exIdInt)
			}
		case float64:
			{
				exIdFloat := exIdRaw.(float64)
				exId = strconv.FormatFloat(exIdFloat, 'f', -1, 64)
			}
		}
		return exId, nil
	}
}
