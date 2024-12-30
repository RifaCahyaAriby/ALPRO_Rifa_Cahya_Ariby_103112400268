package main

import "fmt"

func main() {
	var saldo int
	var transaksi int

	saldo = 0 //  saldo awal dari 0

	for {
		fmt.Print("Masukkan transaksi (0 untuk selesai): ")
		fmt.Scanln(&transaksi)

		if transaksi == 0 { 
			break // keluar dari loop jika input adalah 0
		}

		saldo += transaksi // tambahkan transaksi ke saldo
	}

	fmt.Printf("Saldo akhir: %d\n", saldo) 
}

//PSEUDOCODE
//PROGRAM SaldoAkhir

//DEKLARASI
	//saldo, transaksi : INTEGER

//MULAI
	// Minta input saldo awal
	//TAMPILKAN "Masukkan saldo awal: "
	//BACA saldo

	// Loop untuk meminta input transaksi
	//SELAMA TRUE DO
		// Tampilkan pesan untuk meminta input transaksi
		//TAMPILKAN "Masukkan transaksi (0 untuk selesai): "
		// Baca input transaksi dari pengguna
		// JIKA transaksi = 0 THEN
			// Keluar dari loop jika input adalah 0
		// ENDIF
		// Tambahkan transaksi ke saldo
		//saldo = saldo + transaksi
	//ENDLOOP

	// Tampilkan saldo akhir
	//TAMPILKAN "Saldo akhir: ", saldo DENGAN FORMAT 2 DESIMAL

//SELESAI