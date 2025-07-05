package model

import (
	"log"

	"go-backend-keurani.local/config"
)

type Pendidikan struct {
	ID          int    `json:"id"`
	NamaSekolah string `json:"nama_sekolah"`
	Tahun       int    `json:"tahun"`
}

// Ambil semua data pendidikan
func GetAllPendidikan() ([]Pendidikan, error) {
	db := config.GetDB()
	rows, err := db.Query("SELECT id, nama_sekolah, tahun FROM pendidikan ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pendidikanList []Pendidikan
	for rows.Next() {
		var p Pendidikan
		if err := rows.Scan(&p.ID, &p.NamaSekolah, &p.Tahun); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		pendidikanList = append(pendidikanList, p)
	}
	return pendidikanList, nil
}

// Tambah data pendidikan
func CreatePendidikan(p Pendidikan) error {
	db := config.GetDB()
	_, err := db.Exec("INSERT INTO pendidikan (nama_sekolah, tahun) VALUES ($1, $2)", p.NamaSekolah, p.Tahun)
	return err
}

// Hapus data pendidikan berdasarkan ID
func DeletePendidikan(id int) error {
	db := config.GetDB()
	_, err := db.Exec("DELETE FROM pendidikan WHERE id = $1", id)
	return err
}
