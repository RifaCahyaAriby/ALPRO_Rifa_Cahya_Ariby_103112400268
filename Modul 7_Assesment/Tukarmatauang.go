package main

import "fmt"

func main() {

	// Masukan mata uanggg
	var qirat, dirham, dinar, sisa, fals int

	fmt.Print("Masukan matauang dalam satuan qirat: ")
	fmt.Scan(&qirat)

	// ini merupakan perbandingan dari mata uang
	dinar = qirat / (10 * 10 * 6)
	sisa = qirat % (10 * 10 * 6)

	dirham = sisa / (10 * 6)
	sisa = sisa % (10 * 6)

	fals = sisa / 6
	sisa = sisa % 6

	qirat = sisa

	// menampilkan hasil perbandingan mata uang
	fmt.Println("Hasilnya adalah: ", dinar, dirham, fals, qirat)

}
