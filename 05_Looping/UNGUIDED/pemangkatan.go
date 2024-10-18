
package main

import "fmt"

func main() {
	var a, b, rumus int

	fmt.Print("Masukan bilangan pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukan bilangan pangkat: ")
	fmt.Scanln(&b)

	rumus = 1
	for i := 0; i < b; i++ {
		rumus = rumus * a
	}

	fmt.Print("hasilnya adalah: ", rumus)
}
