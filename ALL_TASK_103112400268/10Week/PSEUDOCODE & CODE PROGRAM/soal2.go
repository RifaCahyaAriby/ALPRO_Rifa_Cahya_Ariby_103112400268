package main

import "fmt"

func main() {
	var a,b,c,d,e float64

	fmt.Print("Masukan temperatur zat radioaktif: ")
	fmt.Scan(&a,&b,&c,&d,&e)

	if a == b && b == c && c == d && d == e {
		fmt.Println("Stabil")
	} else if (a < b && b < c && c < d && d < e) || (a > b && b > c && c > d && d > e) {
		fmt.Println("Stabil Naik/Turun")
	} else {
		fmt.Println("Tidak Stabil")
	}
}

//PSEUDOCODE
//PROGRAM CekStabil

//DEKLARASI
	//a, b, c, d, e : REAL

//MULAI
	// Minta input temperatur zat radioaktif		
	//TAMPILKAN "Masukan temperatur zat radioaktif: "						
	//BACA a, b, c, d, e

	// Cek apakah temperatur zat radioaktif stabil
	//JIKA a == b AND b == c AND c == d AND d == e THEN
		//TAMPILKAN "Stabil"
	//ELSEIF (a < b AND b < c AND c < d AND d < e) OR (a > b AND b > c AND c > d AND d > e) THEN
		//TAMPILKAN "Stabil Naik/Turun"
	//ELSE
		//TAMPILKAN "Tidak Stabil"
	//ENDIF

//SELESAI