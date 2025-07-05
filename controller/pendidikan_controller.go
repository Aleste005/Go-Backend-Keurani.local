package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"go-backend-keurani.local/model"
)

// ----------- API Handler ----------

func GetPendidikan(w http.ResponseWriter, r *http.Request) {
	data, err := model.GetAllPendidikan()
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func CreatePendidikan(w http.ResponseWriter, r *http.Request) {
	var p model.Pendidikan
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	if err := model.CreatePendidikan(p); err != nil {
		http.Error(w, "Gagal menyimpan data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Berhasil disimpan"})
}

func DeletePendidikan(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	if err := model.DeletePendidikan(id); err != nil {
		http.Error(w, "Gagal menghapus data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Berhasil dihapus"})
}

// ----------- View (HTML) Handler ----------

func TampilkanHalamanPendidikan(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/pendidikan_view.gohtml")
	if err != nil {
		http.Error(w, "Gagal memuat template", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Judul": "Daftar Pendidikan",
	}
	tmpl.Execute(w, data)
}

// ----------- Router Registrasi ----------

func RegisterPendidikanRoutes(r *mux.Router) {
	r.HandleFunc("/pendidikan", TampilkanHalamanPendidikan).Methods("GET")

	// Endpoint API JSON
	r.HandleFunc("/api/pendidikan", GetPendidikan).Methods("GET")
	r.HandleFunc("/api/pendidikan", CreatePendidikan).Methods("POST")
	r.HandleFunc("/api/pendidikan/{id}", DeletePendidikan).Methods("DELETE")
}
