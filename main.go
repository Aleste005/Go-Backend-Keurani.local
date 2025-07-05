// File: main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go-backend-keurani.local/config"
	"go-backend-keurani.local/router"
)

func main() {
	// Muat variabel dari .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Gagal membaca file .env")
	}

	// Inisialisasi koneksi ke database
	config.InisialisasiDatabase()

	// Ambil port aplikasi
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("❌ Variabel APP_PORT tidak ditemukan di .env")
	}

	log.Println("🚀 Server berjalan di port:", port)

	// Jalankan server
	err = http.ListenAndServe(":"+port, router.BuatRouterBaru())
	if err != nil {
		log.Fatal("❌ Gagal menjalankan server:", err)
	}
}
