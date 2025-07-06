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

	// 🔄 Redirect "/" ke dashboard
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/RestAPI-BKA", http.StatusFound)
	}).Methods("GET")

	// 🖥️ Halaman dashboard utama
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

	// 🔓 Halaman HTML yang tidak butuh token (View)
	controller.RegisterPenggunaRoutes(router)     // /pengguna
	controller.RegisterPendidikanRoutes(router)   // /pendidikan

	// 🔐 Grup API yang butuh token (JSON)
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthorizationMiddleware)

	// Daftar endpoint API dalam grup /api
	api.HandleFunc("/pengguna", controller.AmbilDataPenggunaAPI).Methods("GET")
	api.HandleFunc("/pendidikan", controller.GetPendidikan).Methods("GET")
	api.HandleFunc("/pendidikan", controller.CreatePendidikan).Methods("POST")
	api.HandleFunc("/pendidikan/{id}", controller.DeletePendidikan).Methods("DELETE")

	// 🔓 Endpoint login (tanpa middleware token)
	router.HandleFunc("/login", controller.Login).Methods("POST")

	return router
}
