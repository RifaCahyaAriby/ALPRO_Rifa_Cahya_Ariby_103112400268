package main

import "fmt"

func main() {

	// Masukan bilangan n
	var n int
	fmt.Print("Masukan n :")
	fmt.Scan(&n)

	// menampilkan hasil dari perkuadratan perulangan dari 1 sampai n
	fmt.Print("Hasilnya adalah: ")
	for i := 1; i <= n; i++ {
		fmt.Print(i*i, " ")
	}

}
