package utils

import (
	"encoding/json"
	"net/http"
)

func KirimJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func KirimError(w http.ResponseWriter, kode int, pesan string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(kode)
	json.NewEncoder(w).Encode(map[string]string{"error": pesan})
}
