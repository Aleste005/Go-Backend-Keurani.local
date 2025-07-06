package controller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"go-backend-keurani.local/model"
	"go-backend-keurani.local/utils"
)

func TampilkanHalamanPengguna(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/user_view.html")
	if err != nil {
		http.Error(w, "Gagal memuat template", http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"Judul": "Daftar Pengguna",
	}
	tmpl.Execute(w, data)
}

func AmbilDataPenggunaAPI(w http.ResponseWriter, r *http.Request) {
	daftar, err := model.AmbilSemuaPengguna()
	if err != nil {
		utils.KirimError(w, http.StatusInternalServerError, "Gagal ambil data pengguna")
		return
	}

	utils.KirimJSON(w, daftar)
}

func RegisterPenggunaRoutes(r *mux.Router) {
	// Halaman biasa (tanpa token)
	r.HandleFunc("/pengguna", TampilkanHalamanPengguna).Methods("GET")

	// Endpoint API (sudah otomatis kena middleware dari router.go)
	r.HandleFunc("/api/pengguna", AmbilDataPenggunaAPI).Methods("GET")
}
