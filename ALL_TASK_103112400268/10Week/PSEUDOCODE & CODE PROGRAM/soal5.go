package main

import "fmt"

func main() {
	var totalbelanja, totalbayar int
	var buat string
	var kartu, diskon, cashback bool

	fmt.Print("Masukan total belanja: ")
	fmt.Scan(&totalbelanja)
	fmt.Print("Bersedia membuat kartu (true/false) ")
	fmt.Scan(&buat)

	if buat == "true" { // Kartu dibuat
		kartu = true
	}
	if totalbelanja >= 100000 {
		diskon = true
		totalbayar = totalbelanja - (totalbelanja * 10 / 100) // Potongan diskon 10%
	} else {
		totalbayar = totalbelanja
	}
	if totalbayar >= 200000 && kartu {
		cashback = true
		totalbayar -= 75000 // Potongan cashback Rp. 75.000
	}
		fmt.Println("Kartu:", kartu)
		fmt.Println("Diskon:", diskon)
		fmt.Println("Cashback:", cashback)
		fmt.Println("Total Bayar:", totalbayar)
}

//PSEUDOCODE
//PROGRAM HitungTotalBelanja

//DEKLARASI
	//totalbelanja : INTEGER
	//buat : STRING	
	//kartu, diskon, cashback : BOOLEAN

//MULAI
	// Minta input total belanja
	//TAMPILKAN "Masukkan total belanja: "
	//BACA totalbelanja

	// Minta input apakah sedang mengikuti asesmen
	//TAMPILKAN "Apakah sedang mengikuti asesmen CLO? (true/false): "
	//BACA buat

	// Cek apakah sedang mengikuti asesmen
	//JIKA sedangasesmen THEN
		// Hitung diskon 35%
		//diskon = 0.35 * totalbelanja
		
		// Kurangi total belanja dengan diskon
		//totalbelanja = totalbelanja - diskon
	//ENDIF

	// Tampilkan total belanja yang harus dibayar dengan format dua desimal
	//TAMPILKAN "Total belanja yang harus dibayar: ", totalbelanja DENGAN FORMAT 2 DESIMAL

//SELESAI