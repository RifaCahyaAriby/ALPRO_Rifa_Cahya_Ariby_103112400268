//program AMPM
//kamus
//jam12, jam24 : integer
//algoritma
//input(jam24)
//if jam24 mod 12 == 0 then
//jam12 = 12
//else
//jam12 = jam24 mod 12
//endif
//if jam24 < 12 then
//output(jam12, "AM")
//else
//output(jam12, "PM")
//endif
//endprogram


package main

import "fmt"

func main() {
	var jam24 int

	// Input jam dalam format 24 jam
	fmt.Print("Masukkan jam : ")
	fmt.Scan(&jam24)

	var jam12 int

	// Menghitung jam dalam format 12 jam
	if jam24%12 == 0 {
		jam12 = 12
	} else {
		jam12 = jam24 % 12
	}

	// Menentukan AM atau PM
	if jam24 < 12 {
		fmt.Println(jam12 , "AM")
	} else {
		fmt.Println(jam12, "PM")
	}
}
