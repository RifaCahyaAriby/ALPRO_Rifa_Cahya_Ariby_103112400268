package main

import "fmt"

func main () {

	// Masukan terdiri dari 2 bilangan yang ingin di jumlahkan
	var  x, y, total int
	fmt.Print("Masukan 2 bilangan yang ingin dijumlahkan :")
	fmt.Scan(&x, &y)
	
	// ini merupakan perulangan penjumlahan
	for i := x; i <= y; i++ {
		total = total + i
	}

	// menampilkan hasil
	fmt.Println("hasilnya adalah: ", total)
}