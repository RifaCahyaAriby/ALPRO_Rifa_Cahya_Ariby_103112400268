package main

import "fmt"

func main () {

	var tinggi, berat, bmi float64
	fmt.Scan(&berat, &tinggi)
	bmi = berat / (tinggi * tinggi)
	fmt.Printf("%.2f", bmi) 

}