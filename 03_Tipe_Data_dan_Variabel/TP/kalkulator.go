package main

import (
	"fmt"
)

func main() {
	var angka1, angka2 float64
	var operator string

	// Input dari pengguna
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	// Kalkulasi berdasarkan operator yang dipilih
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", angka1, angka2, angka1+angka2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", angka1, angka2, angka1-angka2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", angka1, angka2, angka1*angka2)
	case "/":
		// Cek jika pembagian dengan 0
		if angka2 != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", angka1, angka2, angka1/angka2)
		} else {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
		}
	default:
		fmt.Println("Operator tidak valid. Gunakan +, -, *, atau /.")
	}
}
