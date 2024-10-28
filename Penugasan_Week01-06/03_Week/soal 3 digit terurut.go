package main

import (
	"fmt"
)

func main() {
	//berisi variabel dan tipedata integer
	var input int
	var digit1, digit2, digit3 int
	fmt.Println("Masukkan bilangan bulat positif tiga digit:") // masukan untuk pengguna
	fmt.Scan(&input)

	// perhitungan digit dari bilangan
	digit1 = input / 100           // Digit pertama
	digit2 = (input / 10) % 10      // Digit kedua
	digit3 = input % 10             // Digit ketiga

	// mengetahui apakah terurut mengecil
	if digit1 > digit2 && digit2 > digit3 {
		fmt.Println(true) // Keluaran boolean true
	} else {
		fmt.Println(false) // Keluaran boolean false
	}
}

  

