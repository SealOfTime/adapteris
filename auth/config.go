package auth

type Config struct{
	ReceiveAccessCodeUrl string
	Providers []ProviderConfig
}