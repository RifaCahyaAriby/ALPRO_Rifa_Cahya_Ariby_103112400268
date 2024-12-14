package main 

import "fmt"

func main() {
	var bil int

	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bil)

	for i := 1; i <= bil; i++ {
		if i % 2 != 0 {
			fmt.Print(i, " ")
		}
	}
}