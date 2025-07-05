// File: router/router.go
package router

import (
	"net/http"
	"text/template"

	"go-backend-keurani.local/controller"
	"go-backend-keurani.local/middleware"

	"github.com/gorilla/mux"
)

func BuatRouterBaru() http.Handler {
	router := mux.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.CORS)

	// ✅ Redirect dari "/" ke "/RestAPI-BKA"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/RestAPI-BKA", http.StatusFound)
	}).Methods("GET")

	// ✅ Tampilkan halaman dashboard di "/RestAPI-BKA"
	router.HandleFunc("/RestAPI-BKA", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("view/index.html")
		if err != nil {
			http.Error(w, "Gagal memuat halaman", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"Judul": "📡 REST API Keurani - Dashboard",
		}
		tmpl.Execute(w, data)
	}).Methods("GET")

	// ✅ Register semua endpoint API
	controller.RegisterPenggunaRoutes(router)
	controller.RegisterPendidikanRoutes(router)

	return router
}
