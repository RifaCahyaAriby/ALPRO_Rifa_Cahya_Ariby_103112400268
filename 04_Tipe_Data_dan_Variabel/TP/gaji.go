package main

import (
	"fmt"
)

func main() {
	
	var jamKerjaPerMinggu int
	var upahPerJam int
	var gajiMingguan int
	var gajiBulanan int
	var jamNormal int = 40
	var jamLembur int


	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scan(&jamKerjaPerMinggu)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scan(&upahPerJam)


	if jamKerjaPerMinggu > jamNormal {
		jamLembur = jamKerjaPerMinggu - jamNormal
	} else {
		jamLembur = 0
	}

	// Menghitung gaji mingguan dengan rumus 
	gajiMingguan = (jamNormal * upahPerJam) + (jamLembur * 150 / 100 * upahPerJam)

	// Menghitung gaji bulanan 
	gajiBulanan = gajiMingguan * 4

	// Menampilkan total gaji bulanan
	fmt.Println("Gaji bulanan karyawan adalah: Rp", gajiBulanan)
}
