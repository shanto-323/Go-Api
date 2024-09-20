package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/shanto-323/Go-Api/api"
	"github.com/shanto-323/Go-Api/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalance{}
	var decoder *schema.Decoder = schema.NewDecoder()

	var err error
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

  var token  *tools.CoinDetail
  token = (*database).GetUserCoins(params.Username)

  if token == nil {
    log.Error(err)
    api.InternalErrorHandler(w)
    return
  }
  
	var tokenDetails *tools.CoinDetail
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		TotalCoinBalance: (*tokenDetails).Coins,
		SuccessCode:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
