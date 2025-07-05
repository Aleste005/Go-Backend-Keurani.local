// File: config/db.go
package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // import driver MySQL (digunakan secara implisit)
)

// koneksiDB adalah variabel global untuk menyimpan koneksi database
var koneksiDB *sql.DB

// InisialisasiDatabase membuat koneksi ke database menggunakan konfigurasi dari file .env
func InisialisasiDatabase() {
	var err error

	// Ambil nilai dari environment variable untuk membentuk DSN (Data Source Name)
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") +
		"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")

	// Buka koneksi ke database
	koneksiDB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Gagal membuka koneksi ke database:", err)
	}

	// Cek apakah koneksi berhasil dilakukan
	if err = koneksiDB.Ping(); err != nil {
		log.Fatal("❌ Tidak dapat terhubung ke database:", err)
	}

	log.Println("✅ Koneksi ke database berhasil")
}

// GetDB mengembalikan koneksi database aktif
func GetDB() *sql.DB {
	return koneksiDB
}
