package main

import "fmt"

func main()  {
	var (jam, detik, menit, masukan int)
	fmt.Scan(&masukan)

	jam = masukan/3600
	menit = masukan/60
	detik = jam % menit

	fmt.Println(jam,"jam",menit,"menit", detik,"detik")
}