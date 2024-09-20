package tools

import (
	"time"
)

type mockDb struct {
}

var mockLoginDetails = map[string]LoginDetail{
	"alex": {
		AuthToken: "AL323",
		Username:  "ALEX",
	},
	"shanto": {
		AuthToken: "SH323",
		Username:  "SHANTO",
	},
}


var mockCoinDetails = map[string]CoinDetail{
	"alex": {
    Coins: 100,
		Username:  "ALEX",
	},
	"shanto": {
		Coins: 150,
		Username:  "SHANTO",
	},

}


func (d *mockDb) GetUserLoginDetails(username string) *LoginDetail{
  time.Sleep(time.Second * 1)

  var clientData = LoginDetail{}
  clientData, ok := mockLoginDetails[username]

  if !ok {
    return nil
  }
   
  return &clientData
}



func (d *mockDb) GetUserCoinDetails(username string) *CoinDetail{
  time.Sleep(time.Second * 1)

  var clientData = CoinDetail{}
  clientData, ok := mockCoinDetails[username]

  if !ok {
    return nil
  }
   
  return &clientData
}
