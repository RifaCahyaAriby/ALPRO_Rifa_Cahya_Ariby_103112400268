package main

import "fmt"

func main () {
var r int
	var volume, luas float64

	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	volume = (4.0 / 3.0) * 3.1415926535 * float64(r*r*r)
	luas = 4 * 3.1415926535 * float64(r*r)

	fmt.Print("Bola dengan jejari ", r, " memiliki volume ")
	fmt.Println(volume)
	fmt.Print("dan luas kulit ")
	fmt.Println(luas)

}