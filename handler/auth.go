package handler

import (
	"net/http"
)

func AuthMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" && r.URL.Query()["token"] != nil {
			token = r.URL.Query()["token"][0]
		}
		if token == "" {
			responseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		r.Header.Set("Authorization", token)
		fn(w, r)
	}
}
