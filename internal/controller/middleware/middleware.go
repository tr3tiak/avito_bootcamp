package middleware

import (
	"avito_bootcamp/pkg"
	"net/http"
)

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	var hf http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("authorization")
		err := pkg.ValidateToken(token)
		if err != nil {
			return
		}
		h.ServeHTTP(w, r)
	}
	return hf
}
