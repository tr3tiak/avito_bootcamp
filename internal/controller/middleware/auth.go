package middleware

import (
	"avito_bootcamp/pkg"
	"context"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type Role string

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	var hf http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("auth starting")
		tokenWitBearer := r.Header.Get("Authorization")
		token := strings.TrimPrefix(tokenWitBearer, "Bearer ")
		role, err := pkg.ValidateToken(token)
		if err != nil {
			logrus.Error(err)
			return
		}
		ctx := context.WithValue(r.Context(), Role("role"), role)
		r = r.WithContext(ctx)
		logrus.Info("auth complete")
		h.ServeHTTP(w, r)
	}
	return hf
}
