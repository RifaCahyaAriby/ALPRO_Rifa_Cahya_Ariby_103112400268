package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan non negatif: ")
	fmt.Scan(&n)

	faktorial := 1

	for i := 2; i <= n; i++ {
		faktorial = faktorial * i
	}
	fmt.Println("Hasil faktorialnya adalah",faktorial)
}