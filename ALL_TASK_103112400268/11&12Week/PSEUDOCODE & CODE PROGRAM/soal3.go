package main

import	"fmt"

func main() {
	var bil int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&bil)

	total := 0 // vAriabel untuk menyimpan total
	digits := "" // string untuk menyimpan digit

	for bil > 0 {
		digit := bil % 10	// ambil digit terakhir
		total += digit // tambahkan ke total
		digits = fmt.Sprintf("%d %s", digit, digits) // tambahkan ke string digits
		bil /= 10                  // hapus digit terakhir
	}

	fmt.Println(digits) 
	fmt.Println(total)   
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