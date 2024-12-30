package main

import "fmt"
func main() {
	
	var totalbelanja float64
	var sedangasesmen bool

	
	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalbelanja)
	fmt.Print("Apakah sedang mengikuti asesmen CLO? (true/false): ")
	fmt.Scan(&sedangasesmen)


	if sedangasesmen {
		diskon := 0.35 * totalbelanja // Hitung diskon 35%
		totalbelanja -= diskon        // Kurangi total belanja dengan diskon
	}

	
	fmt.Printf("Total belanja yang harus dibayar: %.2f", totalbelanja)
}

//PSEUDOCODE
//PROGRAM HitungTotalBelanja

//DEKLARASI
    //totalbelanja : FLOAT
    //sedangasesmen : BOOLEAN

//MULAI
    // Minta input total belanja
    //TAMPILKAN "Masukkan total belanja: "
    //BACA totalbelanja

    // Minta input apakah sedang mengikuti asesmen
    //TAMPILKAN "Apakah sedang mengikuti asesmen CLO? (true/false): "
    //BACA sedangasesmen

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
