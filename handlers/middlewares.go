package handlers

import (
	"net/http"
)

func AllowHTTPMethodMiddleware(h http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			writeErrorMessage(rw, http.StatusMethodNotAllowed, `HTTP method is not allowed`)
			return
		}

		h.ServeHTTP(rw, req)
	})
}
