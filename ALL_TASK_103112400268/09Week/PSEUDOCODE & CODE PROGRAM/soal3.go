package main

import "fmt"

func main() {
	var huruf string
	fmt.Scan(&huruf)

	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" || 
		huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("bukan konsonan")
	} 
	
	if !(huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" ||
		huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O") {
		fmt.Println("konsonan")
	}
}

//PSEUDOCODE
//PROGRAM CekHuruf

//DEKLARASI
    //huruf : STRING

//MULAI
    // Minta input huruf dari pengguna
    //BACA huruf

    // Cek apakah huruf adalah vokal
    //JIKA huruf == "a" ATAU huruf == "i" ATAU huruf == "u" ATAU huruf == "e" ATAU huruf == "o" ATAU
       //huruf == "A" ATAU huruf == "I" ATAU huruf == "U" ATAU huruf == "E" ATAU huruf == "O" THEN
        //TAMPILKAN "bukan konsonan"
    //ENDIF

    // Cek apakah huruf bukan vokal (dengan negasi)
    //JIKA !(huruf == "a" ATAU huruf == "i" ATAU huruf == "u" ATAU huruf == "e" ATAU huruf == "o" ATAU
           //huruf == "A" ATAU huruf == "I" ATAU huruf == "U" ATAU huruf == "E" ATAU huruf == "O") THEN
        //TAMPILKAN "konsonan"
    //ENDIF

//SELESAI
