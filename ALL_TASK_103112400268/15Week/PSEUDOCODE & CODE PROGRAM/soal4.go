//PSEUDECODE
//program marketDay
//kamus
//algoritma
//input (x, y, z)
// diskon = z * x /100
//if diskon > y then
// diskon = y
//endif
//belanja = z - diskon
//output(belanja, "rupiah")
//endprogram

package main

import "fmt"
func main() {
	var x, y, z float64

	// Input harga barang (x), batas diskon (y), dan jumlah belanja (z)

	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)

	// Menghitung diskon
	diskon := z * x / 100

	// Memastikan diskon tidak melebihi batas
	if diskon > y {
		diskon = y
	}

	// Menghitung total belanja setelah diskon
	belanja := z - diskon

	// Output total belanja
	fmt.Println(belanja,"rupiah")
}
