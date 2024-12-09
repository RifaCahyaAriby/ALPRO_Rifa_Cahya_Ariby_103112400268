package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Scan(&n)

	for i := 1; i <= n/2; i++ { //menentukan bilangan sempurna
		if n%i == 0 { 
			jumlah += i
		}
	}
	if jumlah == n {
		fmt.Println("Ya")
	} else {
		fmt.Println("Bukan")
	}
}
