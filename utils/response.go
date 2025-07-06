package utils

import (
	"encoding/json"
	"net/http"
)

// Fungsi untuk mengirim data JSON
func KirimJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Fungsi untuk mengirim error dalam bentuk JSON
func KirimError(w http.ResponseWriter, kode int, pesan string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(kode)
	json.NewEncoder(w).Encode(map[string]string{"error": pesan})
}

// Fungsi untuk mengirim token login
func KirimToken(w http.ResponseWriter, token string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Fungsi untuk mengirim pesan sukses
func KirimSukses(w http.ResponseWriter, pesan string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": pesan})
}
