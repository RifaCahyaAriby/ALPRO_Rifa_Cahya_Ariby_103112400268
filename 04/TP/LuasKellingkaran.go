package main

import "fmt"

func main() {

	var r float64
	fmt.Print("Masukan jari jari :")
	fmt.Scan(&r)

	// rumus
	Luas := 3.14 * r * r
	Keliling := 2 * 3.14 * r

	fmt.Println("Luasnya adalah ", Luas )
	fmt.Println("Kelilingnya adalah ", Keliling )

}
