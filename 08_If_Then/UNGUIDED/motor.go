package main

import "fmt"

func main() {
	var motor, orang int
	fmt.Scan(&orang)

	motor = orang / 2
	if orang %2 != 0 {
		motor++
	}

	fmt.Println(motor)
}