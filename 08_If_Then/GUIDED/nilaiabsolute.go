package main

import "fmt"

func main() {
	var bil int
	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bil)

	if bil < 0 {
		bil = -bil
	}
	fmt.Print("Nilai absolute :" ,bil)
}