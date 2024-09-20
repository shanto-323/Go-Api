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
