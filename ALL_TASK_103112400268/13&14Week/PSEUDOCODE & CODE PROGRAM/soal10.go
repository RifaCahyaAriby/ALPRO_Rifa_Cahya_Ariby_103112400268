package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan ukuran (bilangan ganjil): ")
    fmt.Scanln(&n)

    if n%2 == 0 {
        fmt.Println("Ukuran harus bilangan ganjil.")
        return
    }

	for j := 1; j <= n; j++ { // menghitung baris
		for i := 1; i <= n; i++ { // menghitung kolom
			if j == i || j == n-i+1 { // menampilkan angka
				fmt.Printf("%d ", j) 
			} else {
				fmt.Print("  ") // menampilkan spasi
			}
		}
		fmt.Println() // spasi
    }
}

//PSEUDOCODE

// MASUKKAN n (ukuran)
//TAMPILKAN "Masukkan ukuran (bilangan ganjil): "
//BACA n

// JIKA n MOD 2 != 0 THEN
	//TAMPILKAN "Ukuran harus bilangan ganjil."
	//KELUAR
//ENDIF

// LOOP UNTUK MENGHITUNG BARIS
//SELAMA j <= n DO
	// LOOP UNTUK MENGHITUNG KOLOM
	//SELAMA i <= n DO
		// JIKA j == i ATAU j == n-i+1 THEN
			//TAMPILKAN j
		// ELSE
			//TAMPILKAN "  "
		//ENDIF
	//ENDLOOP
	//TAMPILKAN "\n"
//ENDLOOP

//SELESAI