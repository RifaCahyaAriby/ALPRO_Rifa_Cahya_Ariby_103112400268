package main 

import "fmt" 

func main() {
	var gol1, gol2, gol3,gol4 int
	var terbanyak , sedikit int

	fmt.Print("masukan 4 gol berturut turut : ") //input
	fmt.Scan(&gol1, &gol2, &gol3, &gol4)

	//mencari gol terbanyak
	if gol1 >= gol2 {
		terbanyak = gol1
	}

	if gol2 > terbanyak {
		terbanyak = gol2
	}

	if gol3 > terbanyak {
		terbanyak = gol3
	}

	if gol4 > terbanyak {
		terbanyak = gol4
	}

	//mencari gol sedikit
	if gol1 <= gol2 {
		sedikit = gol1
	}

	if gol2 < sedikit {
		sedikit = gol2
	}

	if gol3 < sedikit {
		sedikit = gol3
	}

	if gol4 < sedikit {
		sedikit = gol4
	}

	fmt.Println(terbanyak, sedikit)
}


//PSEUDOCODE
// DEKLARASI
//gol1, gol2, gol3, gol4 : INTEGER
//terbanyak, sedikit : INTEGER

// MULAI
// Minta input 4 gol berturut-turut
//TAMPILKAN "Masukkan 4 gol berturut-turut: "
//BACA gol1, gol2, gol3, gol4

// Mencari gol terbanyak
//terbanyak = gol1
//JIKA gol2 > terbanyak THEN terbanyak = gol2 ENDIF
//JIKA gol3 > terbanyak THEN terbanyak = gol3 ENDIF
//JIKA gol4 > terbanyak THEN terbanyak = gol4 ENDIF

// Mencari gol sedikit
//sedikit = gol1
//JIKA gol2 < sedikit THEN sedikit = gol2 ENDIF
//JIKA gol3 < sedikit THEN sedikit = gol3 ENDIF
//JIKA gol4 < sedikit THEN sedikit = gol4 ENDIF

// Tampilkan hasil
//TAMPILKAN "Gol terbanyak:", terbanyak
//TAMPILKAN "Gol sedikit:", sedikit

// SELESAI
