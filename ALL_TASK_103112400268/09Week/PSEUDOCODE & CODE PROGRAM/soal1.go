// #1 CODE GO
package main 

import "fmt"

func main() {
	var penumpang , kapasitasmobil, jumlahpenumpang int

	fmt.Print("Masukkan jumlah penumpang: ")
	fmt.Scan(&penumpang)

	kapasitasmobil = 7

	if penumpang > 0 {
		jumlahpenumpang = penumpang - kapasitasmobil
		fmt.Print(jumlahpenumpang)
	}

}


//PSEUDOCODE
// PROGRAM HitungJumlahPenumpang

//DEKLARASI
    //penumpang : INTEGER
    //kapasitasmobil : INTEGER
    //jumlahpenumpang : INTEGER

//MULAI
    // Minta input jumlah penumpang
    //TAMPILKAN "Masukkan jumlah penumpang: "
    //BACA penumpang

    // Tetapkan kapasitas mobil
    //kapasitasmobil = 7

    // Cek apakah jumlah penumpang lebih dari 0
    //JIKA penumpang > 0 THEN
        // Hitung sisa penumpang setelah mobil terisi penuh
       // jumlahpenumpang = penumpang - kapasitasmobil
        
        // Tampilkan jumlah penumpang yang tidak terangkut
       // TAMPILKAN jumlahpenumpang
   // ENDIF

//SELESAI
