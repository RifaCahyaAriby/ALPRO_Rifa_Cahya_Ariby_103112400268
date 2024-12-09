package main 

import "fmt"

func main() {
	var target, donatur, total, totaldonatur int

	fmt.Print("Masukkan target donasi: ") 
	fmt.Scan(&target)

	for selesai := false; !selesai; {
		fmt.Print("Masukkan jumlah donasi: ") 
		fmt.Scan(&donatur)

		total += donatur // total akan di ambah dari donator
		selesai = (total >= target) // selesai bernilai true jika total >= target

		totaldonatur++ // totaldonatur akan di increment
		fmt.Println("Donatur :", totaldonatur, "Menyumbang", donatur, "total terkumpul:", total)
	}
	fmt.Println("Target tercapai! Total donasi:", total, "dari", totaldonatur, "donatur.")
}

