package main

import "fmt"

func main() {
	// variabel dan tipe data
	var password string 
	var correctPassword = "Rifaganteng" // password yang benar

	// perulangan untuk login sebanyak 3 kali
	for i := 1; i <= 3; i++ {
		fmt.Print("Masukkan password sebelum masuk: ") // untuk meminta input password
		fmt.Scan(&password)

		// Memeriksa apakah password yang dimasukkan benar
		if password == correctPassword {
			fmt.Println("Yeay, login berhasil!") // jika berhasil memasukkan password
			break // keluar dari perulangan setelah login berhasil
		} else {
			fmt.Println("Password salah, coba lagi!") // jika salah memasukkan password
		}

		// Menampilkan pesan setelah percobaan ketiga
		if i == 3 {
			fmt.Println("Login ditolak.") // jika salah memasukkan password sebanyak 3 kali otomatis program akan berhenti
		}
	}
}
