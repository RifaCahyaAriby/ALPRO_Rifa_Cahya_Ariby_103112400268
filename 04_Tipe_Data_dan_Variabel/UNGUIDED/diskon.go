package main

import "fmt"

func main() {

	var belanjaawl, diskon, belanjaahr int

	fmt.Print("total belanja awal: ")
	fmt.Scan(&belanjaawl)
	fmt.Print("diskon: ")
	fmt.Scan(&diskon)

	belanjaahr = belanjaawl - (belanjaawl * diskon / 100)
	fmt.Println("jadi total belanja akhir nya adalah", belanjaahr)

}
