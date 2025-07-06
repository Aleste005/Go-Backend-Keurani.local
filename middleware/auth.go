// File: middleware/auth.go
package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"go-backend-keurani.local/utils"
)

var SecretKey = []byte("rahasia_super_kuat") // ganti sesuai kebutuhan

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.KirimError(w, http.StatusUnauthorized, "Token tidak ditemukan atau format salah")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			utils.KirimError(w, http.StatusUnauthorized, "Token tidak valid")
			return
		}

		// Token valid, lanjut ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
