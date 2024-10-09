package main

import "fmt"

func main() {

	var d1, d2, d3, bilangan int

	fmt.Print("masukan bilangan 3 digit: ")
	fmt.Scan(&bilangan)

	d1 = bilangan / 100
	d2 = (bilangan % 100) / 10
	d3 = bilangan % 10

	fmt.Println(d1 <= d2 && d2 <= d3)

	

}