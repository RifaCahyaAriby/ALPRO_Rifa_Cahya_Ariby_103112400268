package main

import "fmt"

func main () {
	var harga int
	total := 0

	for {
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai) : ")
		fmt.Scan(&harga)
		
		if harga == 0 {
			break
		}
		total += harga // menambahkan semua harga selain 0
	}
	fmt.Println("Total belanja anda:",total)
}
