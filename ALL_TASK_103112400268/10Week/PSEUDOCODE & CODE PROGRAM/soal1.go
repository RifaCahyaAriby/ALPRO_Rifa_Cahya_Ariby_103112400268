package main

import "fmt"

func main() {
	var a, b, c0268 int

	fmt.Print("Masukkan sisi-sisi segitiga: ")
	fmt.Scan(&a, &b, &c0268)

	if a+b > c0268 && a+c0268 > b && b+c0268 > a {
		// Cek segitiga siku-siku terlebih dahulu
		if a*a == b*b+c0268*c0268 || b*b == a*a+c0268*c0268 || c0268*c0268 == a*a+b*b {
			fmt.Println("Segitiga siku-siku")
		} else if a == b && b == c0268 {
			// Segitiga sama sisi
			fmt.Println("Segitiga sama sisi")
		} else if a == b || b == c0268 || a == c0268 {
			// Segitiga sama kaki
			fmt.Println("Segitiga sama kaki")
		} else {
			// Segitiga sembarang
			fmt.Println("Segitiga sembarang")
		}
	} else {
		fmt.Println("Bukan Segitiga")
	}
}

//PSEUDOCODE
//PROGRAM CekSegitiga

//DEKLARASI
	//a, b, c0268 : INTEGER

//MULAI	
	// Minta input sisi-sisi segitiga		
	//TAMPILKAN "Masukkan sisi-sisi segitiga: "						
	//BACA a, b, c0268

	// Cek apakah segitiga siku-siku
	//JIKA a*a == b*b+c0268*c0268 ATAU b*b == a*a+c0268*c0268 ATAU c0268*c0268 == a*a+b*b THEN
		//TAMPILKAN "Segitiga siku-siku"
	//ELSEIF a == b AND b == c0268 THEN
		//TAMPILKAN "Segitiga sama sisi"
	//ELSEIF a == b OR b == c0268 OR a == c0268 THEN
		//TAMPILKAN "Segitiga sama kaki"
	//ELSE
		//TAMPILKAN "Segitiga sembarang"
	//ENDIF

//SELESAI