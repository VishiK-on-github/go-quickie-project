package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234",
		Username:  "alex",
	},
	"jacob": {
		AuthToken: "12345",
		Username:  "jacob",
	},
	"heehoo": {
		AuthToken: "12346",
		Username:  "heehoo",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    1234,
		Username: "alex",
	},
	"jacob": {
		Coins:    12345,
		Username: "jacob",
	},
	"heehoo": {
		Coins:    12346,
		Username: "heehoo",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 2)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 2)

	var coinData = CoinDetails{}
	coinData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &coinData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
