package auth

import (
	"golang.org/x/oauth2"
)

type ProviderConfig struct {
	Name string
	ClientId string
	ClientSecret string
	AuthUrl string
	TokenUrl string
	RedirectUrl string
	Scopes []string
}

func setupProviders(cfg *Config) map[string] oauth2.Config{
	providers := make(map[string] oauth2.Config, len(cfg.Providers))
	for _, p := range cfg.Providers {
		redirectUrl := p.RedirectUrl
		if redirectUrl == "" {
			redirectUrl = cfg.ReceiveAccessCodeUrl
		}
		
		providers[p.Name] = oauth2.Config{
			ClientID:     p.ClientId,
			ClientSecret: p.ClientSecret,
			Endpoint:     oauth2.Endpoint{
				AuthURL:   p.AuthUrl,
				TokenURL:  p.TokenUrl,
				AuthStyle: oauth2.AuthStyleAutoDetect,
			},
			RedirectURL:  redirectUrl,
			Scopes:       p.Scopes,
		} 
	}
	return providers
}
