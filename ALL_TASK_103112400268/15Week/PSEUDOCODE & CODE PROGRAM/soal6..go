package main 

import "fmt"

func main() {
	var tarif, potongan, diskon, tarif_awal float64
	var durasi, kelebihan int
	var member string 

	fmt.Scan(&member, &durasi)

	if member == "Gold" { //menentukan member gold atau sebaliknya
		diskon = 0.5
	} else if member == "Silver" { //menentukan member silver atau sebaliknya
		diskon = 0.25
	} else {
		diskon = 0
	}
	if durasi <= 2 {
		tarif_awal =  float64(durasi) * 65000
	} else {
		kelebihan = durasi - 2
		tarif_awal = 2 * 65000 + float64(kelebihan)*20000
	}
	potongan = diskon * tarif_awal
	tarif = tarif_awal - potongan
	fmt.Println("IDR", tarif)
}

//PSEUDOCODE
//program playground
//kamus
//tarif, potongan, diskon, tarif_awal: real
//durasi, kelebihan: integer
//member: string
//algoritma
//input(member, durasi)
//if member = "Gold" then
//diskon = 0.5
//else if member = "Silver" then
//diskon = 0.25
//else
//diskon = 0
//endif
//if durasi <= 2 then
//tarif_awal = durasi * 65000
//else
//kelebihan = durasi - 2
//tarif_awal = 2 * 65000 + kelebihan * 20000
//endif
//potongan = diskon * tarif_awal
//tarif = tarif_awal - potongan
//output(tarif)
//endprogram