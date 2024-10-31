package main

import "fmt"

func main() {
	var x , y int
	fmt.Scan(&x,&y)

	if y%x == 0 {
		fmt.Println("True")
	}
	if y%x != 0 {
		fmt.Println("False")
	}
	if x%y == 0 {
		fmt.Println("True")
	}
	if x%y != 0 {
		fmt.Println("False")
	}
}
