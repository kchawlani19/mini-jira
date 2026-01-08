package middleware

import (
	"net/http"
	"strings"

	"test_mini_jira/utils"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		token, err := utils.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// optional: extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			_ = claims["user_id"] // future use
			_ = claims["role"]
		}

		next(w, r)
	}
}
