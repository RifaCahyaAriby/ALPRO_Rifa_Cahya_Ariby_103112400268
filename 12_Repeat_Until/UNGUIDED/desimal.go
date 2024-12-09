package main

import "fmt"

func main() {

	var angka, done float32

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// Nilai done digunakan untuk menentukan batas akhir perulangan
	done = float32(int(angka)) + 1

	for selesai := false; !selesai; {

		// Perhitungan ini digunakan untuk memperbesar angka
		angka += 0.1

		// Perhitungan ini digunakan untuk memperkecil angka
		angka = float32(int(angka*10+0.5)) / 10

		// Perhitungan ini digunakan untuk memeriksa apakah angka sudah mencapai batas akhir
		if angka >= done {
			fmt.Printf("%.0f\n", angka)
			selesai = true
		} else {
			fmt.Printf("%.1f\n", angka)
		}

	}

}