package api

import (
	"encoding/json"
	"net/http"
)

// Coin CoinBalance Params
type CoinBalance struct {
	Username string
}

// Coin Coin BalanceResponse
type CoinBalanceResponse struct {
	SuccessCode      int
	TotalCoinBalance int64
}

// Coin Coin BalanceResponse
type Error struct {
	SuccessCode int
	Message     string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		SuccessCode: code,
		Message:     message,
	}

  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(code)

  json.NewEncoder(w).Encode(resp)
}

var (
  RequestErrorHandler = func (w http.ResponseWriter, err error)  {
    writeError(w , err.Error(), http.StatusBadRequest)
  }
  InternalErrorHandler = func (w http.ResponseWriter)  {
    writeError(w , "An Error Happened", http.StatusInternalServerError)
  }
)
