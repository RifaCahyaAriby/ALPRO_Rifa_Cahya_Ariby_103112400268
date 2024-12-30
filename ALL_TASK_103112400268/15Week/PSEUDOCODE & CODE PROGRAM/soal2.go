//PSEUDOCODE
//program serviceMobil
//kamus
//datang, jDatang, mDatang, durasi ,jSelesai, mSelesai: integer
//algotirma
//input(jDatang, mDatang, durasi)
//datang = jDatang * 60 + mDatang + durasi
//jSelesai = datang div 60
//mSelesai = datang mod 60
//if jSelesai < 24 or (jSelesai == 20 and mSelesai == 0) then
//output(jSelesai, mSelesai)
//else
//output("silahkan datang kembali besok")
//endif
//endprogram

package main

import "fmt"

func main() {
	var jDatang, mDatang, durasi int

	// Input waktu kedatangan dan durasi
	fmt.Print("Masukkan jam kedatangan (0-23): ")
	fmt.Scan(&jDatang)
	fmt.Print("Masukkan menit kedatangan (0-59): ")
	fmt.Scan(&mDatang)
	fmt.Print("Masukkan durasi layanan dalam menit: ")
	fmt.Scan(&durasi)

    // Menghitung waktu selesai
	datang := jDatang*60 + mDatang + durasi

	// Menghitung jam dan menit selesai
	jSelesai := datang / 60
	mSelesai := datang % 60

	// Memeriksa apakah waktu selesai valid
	if jSelesai < 24 || (jSelesai == 20 && mSelesai == 0) {
		fmt.Printf("Waktu selesai: %02d:%02d\n", jSelesai, mSelesai)
	} else {
		fmt.Println("Silahkan datang kembali besok")
	}
}
