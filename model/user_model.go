package model

// File: Go-Backend-Keurani.local/model/user_model.go
import (
	"go-backend-keurani.local/config"
)

// Struktur data untuk tabel 'user'
type Pengguna struct {
	ID     int    `json:"id_user"`
	Nama   string `json:"nama"`
	Level  string `json:"level"`
	Alamat string `json:"alamat"`
	NoHP   string `json:"no_hp"`
}

// Ambil semua data pengguna dari database
func AmbilSemuaPengguna() ([]Pengguna, error) {
	db := config.GetDB()

	baris, kesalahan := db.Query("SELECT id_user, nama, level, alamat, no_hp FROM user")
	if kesalahan != nil {
		return nil, kesalahan
	}
	defer baris.Close()

	var daftarPengguna []Pengguna

	for baris.Next() {
		var pengguna Pengguna
		kesalahan := baris.Scan(&pengguna.ID, &pengguna.Nama, &pengguna.Level, &pengguna.Alamat, &pengguna.NoHP)
		if kesalahan != nil {
			return nil, kesalahan
		}
		daftarPengguna = append(daftarPengguna, pengguna)
	}

	return daftarPengguna, nil
}
