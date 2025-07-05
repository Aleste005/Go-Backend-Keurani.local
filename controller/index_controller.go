package controller

import (
	"net/http"
	"text/template"
)

func TampilkanHalamanIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		http.Error(w, "Gagal memuat halaman index", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
