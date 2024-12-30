package main 

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

if angka%3 == 0 {
	fmt.Println("Kelipatan 3")
}

if angka%5 == 0 {
	fmt.Println("Kelipatan 5")
}
}

//PSEUDOCODE

//DEKLARASI
	//angka : INTEGER

//MULAI
	// Minta input angka dari pengguna
	//BACA angka

	// Cek apakah angka kelipatan 3
	//JIKA angka MOD 3 == 0 THEN
		//TAMPILKAN "Kelipatan 3"
	//ENDIF

	// Cek apakah angka kelipatan 5
	//JIKA angka MOD 5 == 0 THEN
		//TAMPILKAN "Kelipatan 5"
	//ENDIF

//SELESAI