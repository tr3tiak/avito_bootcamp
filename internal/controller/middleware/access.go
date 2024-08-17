package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func AccessMiddleware(h http.HandlerFunc) http.HandlerFunc {
	response := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("access starting")
		role := r.Context().Value(Role("role"))
		if role == "moderator" {
			logrus.Info("access complete")
			h.ServeHTTP(w, r)
		} else {
			logrus.Error("access denied")
			return
		}

	})

	return response
}
