package main

import "fmt"

func main (){

	var bb,tb,bmi float64
	fmt.Scan(&bb, &tb)

	bmi= bb * tb * tb
	fmt.Printf("%.0f", bmi)


} 