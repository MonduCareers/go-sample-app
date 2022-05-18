package handlers

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"

	"github.com/MonduCareers/go-sample-app/database"
)

type AccountCreateHandler struct {
	db *sql.DB
}

func NewAccountCreateHandler(db *sql.DB) http.Handler {
	h := &AccountCreateHandler{
		db: db,
	}
	return AllowHTTPMethodMiddleware(h, http.MethodPost)
}

type accountCreateResponse struct {
	Id uuid.UUID `json:"id"`
}

func (h *AccountCreateHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	id, err := database.CreateAccount(h.db)
	if err != nil {
		writeInternalError(rw, err)
		return
	}

	res := accountCreateResponse{
		Id: id,
	}
	writeJSON(rw, http.StatusCreated, res)
}
