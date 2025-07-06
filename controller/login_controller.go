// File: controller/login_controller.go
package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go" // pastikan ini diimport
	"go-backend-keurani.local/middleware"
	"go-backend-keurani.local/utils"
)

// Struktur permintaan login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.KirimError(w, http.StatusBadRequest, "Permintaan tidak valid")
		return
	}

	// Login manual sederhana
	if req.Username != "admin" || req.Password != "123456" {
		utils.KirimError(w, http.StatusUnauthorized, "Username atau password salah")
		return
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		utils.KirimError(w, http.StatusInternalServerError, "Gagal membuat token")
		return
	}

	utils.KirimToken(w, tokenString)
}
