package main

import "fmt"

func main() {
	var bil int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bil)

	if bil < 2 {
		fmt.Println(false)
		return
	}

	prima := true
	for i := 2; i*i <= bil; i++ { // Optimasi dengan memeriksa hingga akar kuadrat bil
		if bil%i == 0 {
			prima = false
			break
		}
	}

	fmt.Println(prima)
}
//PSEUDOCODE
//PROGRAM CekPrima

//DEKLARASI
	//bil : INTEGER

//MULAI
	// Minta input bilangan
	//TAMPILKAN "Masukkan bilangan: "
	//BACA bil

	// Cek apakah bilangan prima
	//JIKA bil < 2 THEN
		//TAMPILKAN false
		//KELUAR
	//ENDIF

	// Cek apakah bilangan prima
	//JIKA bil MOD 2 == 0 ATAU bil MOD 3 == 0 ATAU bil MOD 5 == 0 ATAU bil MOD 7 == 0 THEN
		//TAMPILKAN false
		//KELUAR
	//ENDIF

	// Tampilkan hasil
	//TAMPILKAN true

//SELESAI