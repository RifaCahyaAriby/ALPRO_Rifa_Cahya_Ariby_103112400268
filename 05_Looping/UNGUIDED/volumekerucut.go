package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Masukkan n kerucut: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var r, t, rumus float64

		fmt.Printf("Masukkan jari-jari: ")
		fmt.Scan(&r)
		fmt.Printf("Masukkan tinggi: ")
		fmt.Scan(&t)

		rumus = (1.0 / 3.0) * math.Pi * r * r * t

		fmt.Printf("Volume kerucutnya adalah %.5f", rumus)
	}
}