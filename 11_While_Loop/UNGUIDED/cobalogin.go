package main

import "fmt"

func main() {
	var nama, pw, namabenar, pwbenar string
	var coba int

	namabenar= "Admin" // username
	pwbenar="Admin" // password
	
	// percobaan login
	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&nama)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pw)
	
		if nama == namabenar && pw == pwbenar { // jika username dan password benar
			fmt.Printf("%d percobaan gagal login\n", coba) // jika percobaan gagal
			break // keluar dari perulangan

		} else { // jika username dan password salah
		coba++ // percobaan + 1
		}
	}	
}