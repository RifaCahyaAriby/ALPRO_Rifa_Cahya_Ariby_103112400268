package main

import (
	"fmt"
	"math"
)

func main() {
	var Ax, Ay, Bx, By, Cx, Cy float64

	// Masukkan koordinat titik A, B, dan C
	fmt.Scan(&Ax, &Ay)
	fmt.Scan(&Bx, &By)
	fmt.Scan(&Cx, &Cy)

	// Hitung panjang sisi-sisi segitiga
	AB := math.Sqrt(math.Pow(Bx-Ax, 2) + math.Pow(By-Ay, 2))
	BC := math.Sqrt(math.Pow(Cx-Bx, 2) + math.Pow(Cy-By, 2))
	CA := math.Sqrt(math.Pow(Ax-Cx, 2) + math.Pow(Ay-Cy, 2))

	sisiTerpanjang := math.Max(AB, math.Max(BC, CA))

	fmt.Printf("%.2f\n", sisiTerpanjang)
}
