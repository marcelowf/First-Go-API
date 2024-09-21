package handlers

import (
	"encoding/json"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/marcelowf/First-Go-API/api"
	"github.com/marcelowf/First-Go-API/internal/tools"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var erro error
	
	erro = decoder.Decode(&params, r.URL.Query())

	if erro != nil {
		log.Error(erro)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, erro = tools.NewDatabase()
	if erro != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(erro)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse {
		Balance: (*&tokenDetails).Coins,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	erro = json.NewEncoder(w).Encode(response)
	if erro != nil {
		log.Error(erro)
		api.InternalErrorHandler(w)
		return
	}
}