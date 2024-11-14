package main 

import "fmt"

func main() {
	var kode int

	fmt.Println("Menu Restoran Cepat Saji")
	fmt.Println("1. Burger - Rp.25.000")
	fmt.Println("2. Fried Chicken - Rp.20.000")
	fmt.Println("3. French Fries  - Rp.15.000")
	fmt.Println("4. Soft Drink - Rp.10.000")
	fmt.Println("5. Coffee - Rp.15.000")

	fmt.Print("Masukkan kode menu(1-5): ")
	fmt.Scan(&kode)

	switch kode {
	case 1:
		fmt.Println("Anda memilih Burger - Rp.25.000")
	case 2:
		fmt.Println("Anda memilih Fried Chicken - Rp.20.000")
	case 3:
		fmt.Println("Anda memilih French Fries - Rp.15.000")
	case 4:
		fmt.Println("Anda memilih Soft Drink - Rp.10.000")
	case 5:
		fmt.Println("Anda memilih Coffee - Rp.15.000")
	default:
		fmt.Println("Kode menu tidak valid")
	
	}
	
}