package main

import "fmt"

func main() {
	var input int

	fmt.Print("Masukan bilangan positif :")
	fmt.Scanln(&input)

	for input > 0 { // perulangan selama input lebih besar dari 0
		digit := input % 10 // mendapatkan digit akhir
		fmt.Println(digit) 
		input /= 10 // hapus digit akhir
	}
}