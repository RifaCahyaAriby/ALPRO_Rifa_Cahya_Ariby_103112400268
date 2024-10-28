package main

import "fmt"

func main (){
	// berisi variabel , tipe data (float64 dan string)
	var tinggi,berat float64
	var nama string

	// perintah untuk memasukan informasi bmi
	fmt.Print("Nama : ")
	fmt.Scanln(&nama)
	fmt.Print("Berat kg : ")
	fmt.Scanln(&berat)
	fmt.Print("Tinggi cm : ")
	fmt.Scanln(&tinggi)

	fmt.Println() // ini untuk menambah spasi
	bmi := berat / (tinggi * tinggi) //ini adalah perhitungan BMI

	// ini merupakan keluaran dari hasil perhitungan 
	fmt.Println("Informasi BMI: ")
	fmt.Printf("Nama: %s\n", nama) // artinya menampilkan nama dengan tipe data string
	fmt.Printf("Berat: %.2f kg\n", berat ) // menampilkan bilangan koma dengan belakang kg
	fmt.Printf("Tinggi: %.2f cm\n", tinggi) // menampilkan bilangan koma dengan belakang cm
	fmt.Printf("BMI : %.2f\n ", bmi) // menampilkan hasil perhitungan BMI

}