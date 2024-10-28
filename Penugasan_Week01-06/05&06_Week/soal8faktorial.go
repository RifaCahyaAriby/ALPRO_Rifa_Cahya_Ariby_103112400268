package main

import "fmt"

func main (){
	// berisi variabel dan tipedata integer
	var n, hasil int

	// berisi masukan dan menampilkan bilangan n dari pengguna 
	fmt.Print("masukan bilangan yang ingin difaktorkan: ")
	fmt.Scan(&n)

	// merupakan perulangan peroses faktorial dari n sampai 1 
	hasil = 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}

	// menampilkan hasil faktorial
	fmt.Printf("faktor dari %d adalah %d\n", n, hasil)
}