package main 

import "fmt"

func main() {
	var a,b,c,min,max int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&a,&b,&c)

	if a > b { // a lebih besar dari b
		max = a
		min = b
	} else { // b lebih besar dari a
		max = b
		min = a
	}
	if c > max { // c lebih besar dari max
		max = c
	} else if c < min { // c lebih kecil dari min
		min = c
	}
	fmt.Println("Bilangan terbesar adalah: ", max) 
	fmt.Println("Bilangan terkecil adalah: ", min)

}