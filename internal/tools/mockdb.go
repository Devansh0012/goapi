package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		Username:  "alex",
		AuthToken: "1234",
	},
	"john": {
		Username:  "john",
		AuthToken: "5678",
	},
	"devansh": {
		Username:  "devansh",
		AuthToken: "4321",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Username: "alex",
		Coins:    100,
	},
	"john": {
		Username: "john",
		Coins:    200,
	},
	"devansh": {
		Username: "devansh",
		Coins:    300,
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
