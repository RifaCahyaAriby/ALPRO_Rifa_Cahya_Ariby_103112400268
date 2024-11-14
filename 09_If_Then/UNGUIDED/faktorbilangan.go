package main

import "fmt"

func main() {
	var b int
	var Prima string
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	Prima = "true"
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			if i != 1 && i != b {
				Prima = "false"
			}
		}
	}
	fmt.Println()

	// Menampilkan apakah bilangan adalah prima
	fmt.Print("Prima: ", Prima)
}
