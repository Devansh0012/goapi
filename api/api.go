package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	//success code, usually 200
	Code int

	//acoount balance
	Balance int64
}

// error response
type Error struct {
	//error code
	Code int

	//error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	//write error response
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal Server Error", http.StatusInternalServerError)
	}
)
