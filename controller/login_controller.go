package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go-backend-keurani.local/config"
	"go-backend-keurani.local/utils"
)

// Struktur permintaan login
type LoginRequest struct {
	Username string `json:"username"`
	NoHP     string `json:"no_hp"` // ubah dari password ke no_hp
}

// Struktur JWT claims
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 🔐 Secret key JWT
var jwtKey = []byte("rahasia123")

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.KirimError(w, http.StatusBadRequest, "Format request salah")
		return
	}

	db := config.GetDB()
	row := db.QueryRow("SELECT nama FROM user WHERE username = ? AND no_hp = ?", req.Username, req.NoHP)

	var nama string
	if err := row.Scan(&nama); err != nil {
		utils.KirimError(w, http.StatusUnauthorized, "Login gagal: username atau no HP salah")
		return
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		utils.KirimError(w, http.StatusInternalServerError, "Gagal membuat token")
		return
	}

	utils.KirimToken(w, tokenString)
}
