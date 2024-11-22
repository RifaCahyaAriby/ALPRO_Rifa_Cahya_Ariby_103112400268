package main

import "fmt"

func main() {
	//masukan variabel dan tipe data
	var bil int
	var operasi int
	//masukan bilangan
	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bil)
	//kondisi
	switch { 
		case bil % 10 == 0: // bilangan kelipatan 10
			operasi = bil / 10
			fmt.Println("Kategori: Bilangan Kelipatan 10")
			fmt.Print("Hasil pembagian antara ", bil, "/ 10 =", operasi) //menampilkan hasil
		case bil % 5 == 0: // bilangan kelipatan 5
			operasi = bil * bil
			fmt.Println("Kategori: Bilangan Kelipatan 5")
			fmt.Print("Hasil kuadrat antara ",bil ," ^2 =", operasi) //menampilkan hasil
		case bil % 2 == 0: // bilangan genap
			operasi = bil * (bil+1)
			fmt.Println("Kategori: Bilangan Kelipatan Genap")
			fmt.Print("Hasil perkalian antara ", bil, "*", bil+1, "=", operasi) //menampilkan hasil
		default: // bilangan ganjil
			operasi = bil + (bil+1)
			fmt.Println("Kategori: Bilangan Ganjil ")
			fmt.Print("Hasil penjumlahan antara ", bil, "+", bil+1, "=", operasi) //menampilkan hasil

		}
		
	}

