package main

import "fmt"

func main() {
	var kendaraan string
	var durasi, bayar, tarif int

	fmt.Print("Masukan kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	
	if durasi < 1 {
		durasi = 1 }

	switch kendaraan {
	case "Motor" :
		tarif = 2000
	case "Mobil" :
		tarif =  5000
	case "Truk" :
		tarif = 8000
	default :
		fmt.Println("Jenis kendaraan tidak valid")
	}

	bayar = tarif * durasi

	fmt.Printf("Tarif parkir: Rp %d\n", bayar)
}