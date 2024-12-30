package main

import "fmt"

func main() {
	var profit1, profit2 float64

	fmt.Print("Masukan profit bulan ini dan selanjutnya: ")
	fmt.Scan(&profit1, &profit2)

	if profit1 > profit2{
		kondisi := profit1 - profit2
		fmt.Println("Penurunan sebesar", kondisi)
	}else if profit2 > profit1{
		kondisi := profit2 - profit1
		fmt.Println("Peningkatan sebesar", kondisi)
	}else{
		fmt.Println("Tetap")
	}
}

//PSEUDOCODE
//PROGRAM PenurunanPeningkatan

//DEKLARASI
	//profit1, profit2 : REAL

//MULAI
	// Minta input profit bulan ini dan selanjutnya		
	//TAMPILKAN "Masukan profit bulan ini dan selanjutnya: "						
	//BACA profit1, profit2

	// Cek apakah penurunan atau peningkatan
	//JIKA profit1 > profit2 THEN
		//TAMPILKAN "Penurunan sebesar", profit1-profit2
	//JIKA profit2 > profit1 THEN
		//TAMPILKAN "Peningkatan sebesar", profit2-profit1
	//ELSE
		//TAMPILKAN "Tetap"
	//ENDIF

//SELESAI