package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username  string
	AuthToken string
}

type CoinDetails struct {
	Username string
	Coins    int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(string) *LoginDetails
	GetUserCoins(string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}

func GetDatabase() (*DatabaseInterface, error) {
	return NewDatabase()
}
