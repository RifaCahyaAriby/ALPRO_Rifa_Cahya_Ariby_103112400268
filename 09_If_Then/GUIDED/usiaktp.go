package main

import "fmt"

func main() {
	var usia int
	var kk bool
	fmt.Print("Sudah punya kartu keluarga (true/false): ")
	fmt.Scan(&kk)
	fmt.Print("Usia :" )
	fmt.Scan(&usia)

	if usia >= 17 && kk {
		fmt.Print("bisa membuat ktp")
	} else {
		fmt.Print("belum bisa membuat ktp")
	}

}
