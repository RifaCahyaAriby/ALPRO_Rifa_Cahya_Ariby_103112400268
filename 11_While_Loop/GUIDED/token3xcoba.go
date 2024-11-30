package main

import "fmt"

func main() {
	var token string
	maxPercobaan := 3
	percobaan := 0

	for percobaan < maxPercobaan { // percobaan < 3
		fmt.Print("Masukkan password: ")
		fmt.Scan(&token)

		if token == "12345abcde" { // jika token benar
			fmt.Println("Selamat, Anda berhasil login!")
			return
		}

		percobaan++ // percobaan akan ditambah 1 sampai percobaan < 3
		if percobaan < maxPercobaan { 
			fmt.Println("Maaf, password salah") // jika token salah
		}
	}
	fmt.Println("Anda tidak bisa login, kesempatan anda sudah 3x login habis!") // jika percobaan sudah 3x
}
