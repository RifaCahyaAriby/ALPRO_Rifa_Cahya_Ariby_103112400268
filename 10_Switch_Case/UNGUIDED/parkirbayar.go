package main

import "fmt"

func main() {
	// masukan variabel dan tipe data
	var kendaraan string
	var durasi, bayar, tarif int

	// scan masukan kendaraan dan durasi / jam
	fmt.Print("Masukan kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	// menyatakan jika durasi kurang dari 1 maka akan di nyatakan 1 jam
	if durasi < 1 {
		durasi = 1 }

	// kendaraan dengan beberapa aksi 
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
	// perhitungan untuk tarif dan durasi lama parkir /jam
	bayar = tarif * durasi
	// mencetak hasil jumlah tarif parkir
	fmt.Printf("Tarif parkir: Rp %d\n", bayar)
}
