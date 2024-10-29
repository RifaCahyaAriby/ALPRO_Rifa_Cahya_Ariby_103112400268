package main

import "fmt"

func main () {
	var nilai int
	fmt.Print("Nilai ujian anda :")
	fmt.Scan(&nilai)

	fmt.Println()

	if nilai >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}