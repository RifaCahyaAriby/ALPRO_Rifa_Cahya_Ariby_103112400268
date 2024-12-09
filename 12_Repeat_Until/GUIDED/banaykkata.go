package main

import "fmt"

func main() {
	var kata string
	var banyak int

	fmt.Scan(&kata, &banyak)

	counter := 0
	
	for done := false; !done; {
		fmt.Println(kata)
		counter++
		done = (counter >= banyak)
	}
}