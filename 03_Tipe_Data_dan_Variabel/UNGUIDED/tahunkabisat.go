package main

import "fmt"

	func main() {
		var tahun int
		fmt.Print("Tahun: ")
		fmt.Scan(&tahun)
		kabisat := (tahun/400*400 == tahun) || (tahun/4*4 == tahun && tahun/100*100 != tahun)
		fmt.Println("Kabisat:", kabisat)
	}