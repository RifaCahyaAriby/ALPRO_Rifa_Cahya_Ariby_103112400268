package main

import "fmt"

func main () {
	var j, a, b int
	var hasil int

	
	fmt.Print("Masukan bilangan bulat positif: ")
	fmt.Scan(&a,&b)

	for j = 1; j<=b; j++ {
		hasil = hasil + a
	}

	fmt.Println("hasilnya ", hasil)

}