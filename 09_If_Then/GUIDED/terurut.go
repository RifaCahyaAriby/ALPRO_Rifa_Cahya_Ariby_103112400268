package main

import "fmt"

func main() {
	var digit1, digit2, digit3, digit4 int
	var bil int
	var teks string
	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)
	if bil < 1000 {
		fmt.Println("Bilangan harus 4 digit")
		return
	}
	digit1 = bil / 1000
	digit2 = (bil % 1000) / 100
	digit3 = (bil % 100) / 10
	digit4 = bil % 10


	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		teks = "Terurut membesar"
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
		teks = "Terurut mengecil"
	} else {
		teks = "Tidak terurut"
	}
	fmt.Println("digit pada bilangan",bil,teks)
}1