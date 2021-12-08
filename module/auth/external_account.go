package auth

type ExternalAccount struct {
	Id int
	UserId int
	ExternalId string
}

type ExternalAccountRepository interface {
	ReadAll() ([]ExternalAccount, error)
}
