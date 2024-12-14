package main 

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)

	for i := 2; i < bil; i++ { // menentukan bilangan prima
		if bil % i == 0 {
			fmt.Println("bukan prima")
			return
		}
	}
	fmt.Println("prima")
}
