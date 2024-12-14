package main 

import "fmt"

func main () {
	var n int
	var bunga string
	fmt.Print("N: ")
	fmt.Scan(&n)

	pita := ""
	for i := 1; i <= n; i++ {
		fmt.Printf("Bunga %d : ", i)
		fmt.Scan(&bunga)
		if i == 1 {
			pita += bunga
		} else {
			pita += " – " + bunga
		}
	}
	fmt.Println("Pita: ", pita)
}