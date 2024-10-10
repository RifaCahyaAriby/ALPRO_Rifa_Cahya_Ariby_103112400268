package main

import (	
	"fmt"
)

func main()  {
	var fahrenheit float64
	fmt.Println ("Masukan suhu dalam derajat fahrenheit: ")
	fmt.Scan(&fahrenheit)

	kelvin := (fahrenheit-32)*5/9 + 273.15
	fmt.Printf("suhu dalam derajat Kalvin: %.2f K\n", kelvin)

}

