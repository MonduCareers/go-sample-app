package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeErrorMessage(rw http.ResponseWriter, statusCode int, errorMessage string) {
	errorResponse := struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: errorMessage,
	}

	writeJSON(rw, statusCode, errorResponse)
}

func writeJSON(rw http.ResponseWriter, statusCode int, res interface{}) {
	body, err := json.Marshal(res)
	if err != nil {
		writeInternalError(rw, err)
		return
	}

	writeResponse(rw, statusCode, body)
}

func writeInternalError(rw http.ResponseWriter, err error) {
	fmt.Println(err)
	writeResponse(rw, http.StatusInternalServerError, []byte(`{"error_message":"internal error"}`))
}

func writeResponse(rw http.ResponseWriter, statusCode int, body []byte) {
	rw.Header().Set("content-type", "application/json; charset=utf-8")
	rw.WriteHeader(statusCode)
	if _, err := rw.Write(body); err != nil {
		fmt.Println(err)
	}
}
