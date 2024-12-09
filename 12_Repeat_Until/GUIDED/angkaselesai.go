package main 

import "fmt"

func main() {
	var angka int
	var selesai bool
	for selesai = false; !selesai; {
		fmt.Scan(angka)
		
		selesai = angka > 0
	}
	fmt.Printf("%d adl biangan bulat positif\n", angka)
}