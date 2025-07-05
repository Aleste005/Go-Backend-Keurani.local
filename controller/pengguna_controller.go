package controller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"go-backend-keurani.local/model"
	"go-backend-keurani.local/utils"
)

// Fungsi untuk render halaman
func TampilkanHalamanPengguna(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/user_view.gohtml")
	if err != nil {
		http.Error(w, "Gagal memuat template", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Judul": "Daftar Pengguna",
	}
	tmpl.Execute(w, data)
}

// ✅ Fungsi API untuk ambil data pengguna dalam format JSON
func AmbilDataPenggunaAPI(w http.ResponseWriter, r *http.Request) {
	daftar, err := model.AmbilSemuaPengguna()
	if err != nil {
		utils.KirimError(w, http.StatusInternalServerError, "Gagal ambil data pengguna")
		return
	}

	utils.KirimJSON(w, daftar)
}

// ✅ Daftarkan semua route pengguna (view dan API)
func RegisterPenggunaRoutes(r *mux.Router) {
	// View
	r.HandleFunc("/pengguna", TampilkanHalamanPengguna).Methods("GET")

	// API
	r.HandleFunc("/api/pengguna", AmbilDataPenggunaAPI).Methods("GET")
}
