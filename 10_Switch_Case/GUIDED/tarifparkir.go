package main

import "fmt"

func main() {
	// masukan variabel dan tipedata
	var kendaraan string
	var durasi int
	var tarif int

	// scan input kendaraan dan durasi lama parkir
	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch { // kondisi
	case kendaraan == "Motor" && durasi >= 1 && durasi <= 2: 
		tarif = 7000
	case kendaraan == "Motor" && durasi > 2:
		tarif = 9000
	case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2: 
		tarif = 15000
	case kendaraan == "Mobil" && durasi > 2: 
		tarif = 20000
	case kendaraan == "Truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case kendaraan == "Truk" && durasi > 2:
		tarif = 35000
	default: // kondisi default
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
// cetak hasil tarif
	fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}
