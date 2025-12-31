package middleware

import (
	"net/http"
	"strings"

	"mini-jira/auth"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, err := auth.ValidateToken(parts[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
