package main

import "fmt"

func main() {
	
	var jammulai, menitmulai, jamselesai, menitselesai int
	fmt.Scan(&jammulai, &menitmulai, &jamselesai, &menitselesai)

	
	mulaiTotalMenit := jammulai*60 + menitmulai
	selesaiTotalMenit := jamselesai*60 + menitselesai

	
	if selesaiTotalMenit < mulaiTotalMenit {
		
		selesaiTotalMenit += 12 * 60
	}
	durasiMenit := selesaiTotalMenit - mulaiTotalMenit
	jam := durasiMenit / 60
	menit := durasiMenit % 60

	fmt.Printf("%d jam %d menit\n", jam, menit)
}


//PSEUDOCODE
//DEKLARASI
	//jammulai, menitmulai, jamselesai, menitselesai : INTEGER

//MULAI
	// Minta input jam mulai, menit mulai, jam selesai, dan menit selesai
	//BACA jammulai, menitmulai, jamselesai, menitselesai

	// Hitung durasi dalam menit
	//durasiMenit = (jamselesai * 60 + menitselesai) - (jammulai * 60 + menitmulai)

	// Tampilkan durasi dalam jam dan menit
	//TAMPILKAN durasiMenit / 60, durasiMenit % 60

//SELESAI