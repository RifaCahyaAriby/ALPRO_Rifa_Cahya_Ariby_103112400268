package main 

import "fmt"

func main () {
	var x, y, jumlah int
	fmt.Scan(&x, &y)
	// division
	for x >= y { 
		x = x - y
		jumlah++ // jumlah akan menambah setiap kali x kurang dari y
	}
	fmt.Print(jumlah)
}