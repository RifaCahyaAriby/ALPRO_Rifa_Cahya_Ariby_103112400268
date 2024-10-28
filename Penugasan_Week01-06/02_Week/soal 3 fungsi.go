package main // deklarasi

import "fmt"

func main() {
	//berisi variabel dan tipe data  (integer dan float64)
	var x, y int	
	var persamaan float64
	// perintah masukan bilangan dari pengguna 
	fmt.Print("masukan bilangan x y : ")
	fmt.Scan(&x,&y)

	//berisi persamaan yang terdapat pada soal
	persamaan = 1/(3 * float64(x)*float64(x) + 10) + 10 * float64 (y) + 7
	
	fmt.Println("hasilnya adalah :" ,persamaan)

}