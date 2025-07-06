package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("rahasia123")

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, `{"error":"Token tidak ditemukan atau format salah"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, `{"error":"Token tidak valid atau expired"}`, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
