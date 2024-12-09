package main

import "fmt"

func main (){
	var angka int
	var digit int

	fmt.Scan(&angka)

	for selesai := false; !selesai; { // perulangan selama selesai bernilai false
		angka /= 10 // hapus digit akhir

		selesai = angka == 0 // selesai bernilai true jika angka bernilai 0
		digit++ // increment digit
	}
	fmt.Println(digit)
}