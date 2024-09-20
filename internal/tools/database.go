package tools
import (
	log "github.com/sirupsen/logrus"
)

type LoginDetail struct {
	AuthToken string
	Username  string
}
type CoinDetail struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetail
	GetUserCoins(username string) *CoinDetail
	SetUpDataBase() error
}


func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}

	var err error = database.SetUpDataBase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
