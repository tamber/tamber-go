package account

import (
	"github.com/tamber/tamber-go"
)

var (
	AccountUrl = "https://dashboard.tamber.com"
)

type Account struct {
	S         *tamber.SessionConfig
	Email     string
	Password  string
	AuthToken *tamber.AuthToken
}

func GetDefaultAccountSessionConfig() *tamber.SessionConfig {
	config := tamber.GetDefaultSessionConfig()
	config.URL = AccountUrl
	config.HTTPClient.Timeout = 0
	return config
}

func (a *Account) Init(email, pw string, config *tamber.SessionConfig) error {
	if config == nil {
		config = GetDefaultAccountSessionConfig()
	}
	a.Email = email
	a.Password = pw
	authToken, err := a.Login()
	if err != nil {
		return err
	}
	a.AuthToken = authToken
	return nil
}

// New creates a new Account object
func New(email, pw string, config *tamber.SessionConfig) *Account {
	account := Account{}
	account.Init(email, pw, config)
	return &account
}
