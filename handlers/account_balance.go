package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/MonduCareers/go-sample-app/database"
)

type AccountBalanceHandler struct {
	db *sql.DB
}

func NewAccountBalanceHandler(db *sql.DB) http.Handler {
	h := &AccountBalanceHandler{
		db: db,
	}
	return AllowHTTPMethodMiddleware(h, http.MethodGet)
}

type accountBalanceResponse struct {
	Balance int `json:"balance"`
}

func (h *AccountBalanceHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	if len(idStr) == 0 {
		writeErrorMessage(rw, http.StatusBadRequest, `Missing "id" in the query string`)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println(err)
		writeErrorMessage(rw, http.StatusBadRequest, `Invalid format for "id" param. Must in in UUID`)
		return
	}

	balance, err := database.GetAccountBalance(h.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeErrorMessage(rw, http.StatusNotFound, `Account not found`)
			return
		}

		writeInternalError(rw, err)
		return
	}

	res := accountBalanceResponse{
		Balance: balance,
	}
	writeJSON(rw, http.StatusOK, res)
}
