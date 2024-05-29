package main

import (
	"fmt"
)

// Data struct untuk menyimpan informasi kendaraan
type Data struct {
	Nama string
	Plat string
	Tipe string
	Rute []string
}

// Pelanggaran struct untuk menyimpan informasi pelanggaran
type Pelanggaran struct {
	Name   string
	Tilang int
}

// Function kenaRazia untuk memeriksa pelanggaran aturan ganjil-genap
func kenaRazia(tanggal int, data []Data) []Pelanggaran {
	// Lokasi Ganjil Genap
	ruteGanjilGenap := []string{"Gajah Mada", "Hayam Wuruk", "Sisingamangaraja", "Panglima Polim", "Fatmawati", "Tomang Raya"}

	var hasil []Pelanggaran

	for _, kendaraan := range data {
		// Memeriksa jenis kendaraan, hanya MOBIL yang diperiksa
		if kendaraan.Tipe != "Mobil" && kendaraan.Tipe != "mobil" {
			continue
		}

		// Mendapatkan digit terakhir dari nomor plat kendaraan
		var angkaTerakhirChar byte
		for i := len(kendaraan.Plat) - 1; i >= 0; i-- {
			if kendaraan.Plat[i] >= '0' && kendaraan.Plat[i] <= '9' {
				angkaTerakhirChar = kendaraan.Plat[i]
				break
			}
		}

		// Memeriksa apakah nomor plat sesuai dengan aturan ganjil-genap pada tanggal tersebut
		angkaTerakhir := int(angkaTerakhirChar - '0')
		tanggalGanjil := tanggal%2 != 0
		platGanjil := angkaTerakhir%2 != 0
		pelanggaranN := 0




		for _, rute := range kendaraan.Rute {
			for _, ruteGG := range ruteGanjilGenap {
				if rute == ruteGG {
					if tanggalGanjil != platGanjil {
						pelanggaranN++
					}
				}
			}
		}

		if pelanggaranN > 0 {
			hasil = append(hasil, Pelanggaran{
				Name:   kendaraan.Nama,
				Tilang: pelanggaranN,
			})
		}
	}

	return hasil
}
