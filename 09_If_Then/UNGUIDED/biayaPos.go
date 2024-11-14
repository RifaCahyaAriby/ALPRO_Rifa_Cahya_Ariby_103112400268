package main

import "fmt"
func main() {
	var beratgram,sisagram, kg, biayadasar, totalBiaya, biayatambahan int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratgram)

	kg = beratgram / 1000
	sisagram = beratgram % 1000	
	biayadasar = kg * 10000

	if kg > 10 {
		biayatambahan = 0
	} else if sisagram >= 500 {
		biayatambahan = sisagram * 5
	} else {
		biayatambahan = sisagram * 15
	}
	totalBiaya = biayadasar + biayatambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisagram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayadasar, biayatambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)

}
