package main

import "fmt"

func main() {
 var nam float64 
 var nmk string
 fmt.Print("Nilai akhir mata kuliah:  ")
 fmt.Scan(&nam)
 if nam > 80 {
 nmk = "A"
 }  else if 72.5 < nam && nam <= 80 {
	nmk = "AB"
	} else if 65 < nam && nam <= 72.5 {
	nmk = "B"
	} else if 57.5 < nam && nam <= 65 {
	nmk = "BC"
	} else if 50 < nam && nam <= 57.5 {
	nmk = "C"
	} else if 40 < nam && nam <= 50 {
	nmk = "D"
	} else if nam <= 40 {
	nmk = "E"
	}
   
	fmt.Println("Nilai mata kuliah:  ", nmk)
}

// 1. jika saya memasukan nam 80.1 maka akan menghasilkan A
// 2. kesalahan dari program yang terdapat pada modul yaitu
// - kurang nya else untuk melakukan percabangan
// - ada kesalahan penulisan tanda >
// - perlu disesuaikan posisi else yang benar
// - ada kesalahan penulisan variabel nam pada variable nmk
// 3. Sudah saya perbaiki dan hasilnya benar