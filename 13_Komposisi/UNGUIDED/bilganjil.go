package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	n = 1 //menentukan bilangan ganjil
	for i := 1; i <= n; i++ { 
		if i % 2 != 0 {
			n++
		}
	}
	fmt.Println("Terdapat", n, " bilangan ganjil ")
}