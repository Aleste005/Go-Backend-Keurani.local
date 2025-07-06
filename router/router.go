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

	// ✅ Tampilkan halaman dashboard
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

	// 🔐 Protected API group
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthorizationMiddleware) // 🔐 Pasang middleware auth di sini

	// 🔐 Daftarkan semua endpoint yang butuh token di bawah ini
	controller.RegisterPenggunaRoutes(api)
	controller.RegisterPendidikanRoutes(api)

	// (Opsional) Tambahkan endpoint login (tanpa middleware auth)
	router.HandleFunc("/login", controller.Login).Methods("POST")

	return router
}
