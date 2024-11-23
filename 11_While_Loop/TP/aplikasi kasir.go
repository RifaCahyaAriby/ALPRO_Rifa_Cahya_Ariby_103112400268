package main

import "fmt"

func main() {
	// variabel dan tipe data
	var total float64
	var harga, jumlah int
	var lanjutkan string

	// perulangan untuk menambahkan barang
	for {
		fmt.Print("Masukkan harga barang: Rp. ") // untuk meminta input harga barang
		fmt.Scan(&harga)
		fmt.Print("Masukkan jumlah barang: ") // untuk meminta input jumlah barang
		fmt.Scan(&jumlah)
		
		total += float64(harga * jumlah) // menghitung total belanja

		fmt.Print("Apakah Anda ingin menambahkan barang lagi? (ya/tidak): ") 
		fmt.Scan(&lanjutkan) // untuk meminta input apakah ingin menambahkan barang lagi atau tidak

		if lanjutkan == "tidak" { 
			break // jika tidak ingin menambahkan barang lagi, keluar dari perulangan
		}
	}

	fmt.Printf("Total belanja: Rp. %.2f\n", total) // menampilkan total belanja
	fmt.Println("Transaksi selesai.") // menyelesaikan transaksi
}
