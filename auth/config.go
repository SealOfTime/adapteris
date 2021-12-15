package auth

import (
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/url"
	"strings"
)

type Config struct{
	ReceiveAccessCodeUrl string
	Providers []ProviderConfig
}

type ProviderConfig struct {
	oauth2.Config
	Name string
}

func setupProviders(cfg *Config) []Authenticator{
	providers := make([]Authenticator, len(cfg.Providers))
	for _, p := range cfg.Providers {
		redirectUrl := p.RedirectURL
		if redirectUrl == "" {
			redirectUrl = fmt.Sprintf("%s/%s", cfg.ReceiveAccessCodeUrl, strings.ToLower(p.Name))
		}

		if _, err := url.Parse(redirectUrl); err != nil {
			log.Fatalf("Couldn't initialize authentication provider %s with redirectUrl '%s': %+v", p.Name, redirectUrl, err)
		}
		//todo: actually initialise providers separate into WellKnownProviders with specific OAuth2 token handling and UnknownProvider with unified process
	}
	return providers
}