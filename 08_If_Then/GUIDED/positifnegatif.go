package main 

import "fmt"

func main() {
	var bil int 
	fmt.Scan(&bil)

	if bil > 0 {
		fmt.Println("positif")
	} else {
			fmt.Println("bukan positif")
		}
	}
