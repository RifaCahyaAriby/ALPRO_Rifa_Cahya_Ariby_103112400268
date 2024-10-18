package main

import "fmt"

func main() {
	var n, rumus int

	fmt.Print("Masukan bilangan positif: ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		rumus = rumus + i
	}
	fmt.Print("hasilnya yaitu: ", rumus)
}