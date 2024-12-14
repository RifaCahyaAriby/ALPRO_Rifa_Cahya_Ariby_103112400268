package main 

import "fmt"

func main() {
	var percobaan string

	urutan := "merah kuning hijau ungu"
	berhasil := true 

	for i := 1; i <= 5; i++ { // menampilkan percobaan
		fmt.Printf("Percobaan %d :", i)
		fmt.Scan(&percobaan)
		berhasil = berhasil && (percobaan == urutan)
		}
		fmt.Println("berhasil =", berhasil)
	}
