package main

import "fmt"

func main() {
	
	var n, m, x, y int
	fmt.Print("Masukkan n (jumlah gula), m (jumlah kopi), x (gula per cangkir), y (kopi per cangkir): ")
	fmt.Scanln(&n, &m, &x, &y)
	cangkir := 0 // Inisialisasi variabel untuk menyimpan jumlah cangkir

	for n >= x && m >= y {
		n -= x // Kurangi gula
		m -= y // Kurangi kopi
		cangkir++
}
fmt.Println(cangkir)
}

//PSEUDOCODE

// MASUKKAN n (jumlah gula), m (jumlah kopi), x (gula per cangkir), y (kopi per cangkir)
//TAMPILKAN "Masukkan n (jumlah gula), m (jumlah kopi), x (gula per cangkir), y (kopi per cangkir): "
//BACA n, m, x, y

// Inisialisasi variabel cangkir untuk menyimpan jumlah cangkir
//cangkir = 0

// Loop selama n >= x dan m >= y
//SELAMA n >= x DAN m >= y DO
	// Kurangi gula dan kopi
	//n = n - x
	//m = m - y
	// Tambahkan 1 ke cangkir
	//cangkir = cangkir + 1
//ENDLOOP

//TAMPILKAN cangkir

//SELESAI