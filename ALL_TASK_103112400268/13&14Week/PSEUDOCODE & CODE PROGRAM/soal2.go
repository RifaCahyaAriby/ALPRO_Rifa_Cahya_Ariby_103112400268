package main

import (
	"fmt"
)

func main() {
	var decimal int
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&decimal)

	if decimal == 0 {
		return
	}

	binary := "" // Inisialisasi variabel untuk menyimpan hasil konversi
	for n := decimal; n > 0; n /= 2 { // untuk melakukan konversi
		binary = fmt.Sprintf("%d%s", n%2, binary) // Menambahkan digit biner ke string
	}

	fmt.Print( binary)
}
//PSEUDOCODE

//DEKLARASI
	//decimal : INTEGER
	//binary : STRING

//MULAI
	// Minta input bilangan desimal
	//TAMPILKAN "Masukkan bilangan desimal: "
	//BACA decimal

	// Konversi bilangan desimal ke biner
	//binary = ""	// Inisialisasi variabel untuk menyimpan bilangan biner
	//SELAMA decimal > 0 DO
		//binary = fmt.Sprintf("%d%s", decimal%2, binary) // Menambahkan digit biner ke string
		//decimal /= 2	// Hapus digit terakhir
	//ENDLOOP

	// Tampilkan hasil konversi
	//TAMPILKAN binary

//SELESAI