package main

import "fmt"

func main () {
	var a, b int
	fmt.Print("Masukan bilangan a: ")
	fmt.Scan(&a)
	fmt.Print("Masukan bilangan b: ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		fmt.Println(" ", i)
    }
}