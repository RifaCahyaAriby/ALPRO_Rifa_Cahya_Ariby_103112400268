package main 

import "fmt"

func main() {
	var usia int

	fmt.Print("Masukan usia anda: ")
	fmt.Scan(&usia)

	switch {
	case usia >= 0 && usia <= 12:
		fmt.Println("Anak-anak")
	case usia >= 13 && usia <= 17:
		fmt.Println("Remaja")
	case usia >= 18 && usia <= 64:
		fmt.Println("Dewasa")
	case usia >= 65:
		fmt.Println("Lansia")
	default:
		fmt.Println("Usia tidak valid")
	}

	}