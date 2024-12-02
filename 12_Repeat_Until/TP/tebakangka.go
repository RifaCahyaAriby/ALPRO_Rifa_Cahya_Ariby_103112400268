package main

import "fmt"

func main() {

	var tebakan int
	angkarahasia := 7 

	for {
		fmt.Print("Tebak angka (1-10): ") 
		fmt.Scan(&tebakan)
	
		if tebakan < 1 || tebakan > 10 { // Memeriksa apakah tebakan antara 1 dan 10
			fmt.Println("Angka harus antara 1 hingga 10!")
			continue
		}

		if tebakan != angkarahasia { // Memeriksa apakah tebakan benar
			fmt.Println("Tebakan Anda salah, coba lagi.")
			continue
		}
		fmt.Println("Selamat, tebakan Anda benar!")
		break // berhenti
	}
}