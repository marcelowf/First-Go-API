package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Marcelo": {
		AuthToken: "123ABC",
		Username:  "Marcelo",
	},
	"Laura": {
		AuthToken: "456DEF",
		Username:  "Laura",
	},
	"Aurora": {
		AuthToken: "789HIJ",
		Username:  "Aurora",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Marcelo": {
		Coins:    1000000,
		Username: "Marcelo",
	},
	"Laura": {
		Coins:    500000,
		Username: "Laura",
	},
	"Aurora": {
		Coins:    10,
		Username: "Aurora",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

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
