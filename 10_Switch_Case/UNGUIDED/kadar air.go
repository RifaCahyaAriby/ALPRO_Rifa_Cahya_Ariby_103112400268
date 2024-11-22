package main

import "fmt"

func main (){
	// berisi variabel dan tipe data
	var ph float64
	fmt.Print("Masukan pH air: ") // scan masukan pH
	fmt.Scan(&ph)
	// kondisi
	switch {
	case ph < 0 || ph > 14: // akan memeriksa apakah ph kurang dari 0 atau lebih besar dari 14
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	case ph >= 6.5 && ph <= 8.6: // akan memeriksa apakah ph kurang dari 6.5 atau lebih besar dari 8.6
		fmt.Println("Air layak diminum")
	default: // akan memeriksa apakah ph tidak memenuhi kondisi sebelumnya
		fmt.Println("Air tidak layak diminum")
	}
}