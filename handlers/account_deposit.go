package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/MonduCareers/go-sample-app/database"
)

type AccountDepositHandler struct {
	db *sql.DB
}

func NewAccountDepositHandler(db *sql.DB) http.Handler {
	h := &AccountDepositHandler{
		db: db,
	}
	return AllowHTTPMethodMiddleware(h, http.MethodPost)
}

type accountDepositRequest struct {
	Id     uuid.UUID `json:"id"`
	Amount int       `json:"amount"`
}

func (h *AccountDepositHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var body accountDepositRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		writeErrorMessage(rw, http.StatusBadRequest, "Request contains invalid JSON payload")
		return
	}

	if body.Amount < 1 {
		writeErrorMessage(rw, http.StatusUnprocessableEntity, `Param "amount" must be greater than 0`)
		return
	}

	err = database.DepositAmount(h.db, body.Id, body.Amount)
	if err != nil {
		writeInternalError(rw, err)
		return
	}

	writeResponse(rw, http.StatusNoContent, nil)
}
