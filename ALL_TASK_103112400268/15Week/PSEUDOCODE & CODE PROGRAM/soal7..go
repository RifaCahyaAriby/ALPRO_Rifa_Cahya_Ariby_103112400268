package main

import "fmt"

func main() {
	var x, j, bilangan, nX, nZero int

	fmt.Scan(&x)
	for j = 1; j <= 9; j++ {
		fmt.Scan(&bilangan)
		if bilangan == x {
			nX++
		} else {
			nZero++
		}
	}

	if nX >= nZero {
		fmt.Printf("Modus = %d\n", x)
	} else {
		fmt.Printf("Modus = 0\n")
	}
}


//PSEUDOCODE
//program modus
//kamus
//x, j, bilangan, nX, nZero: integer
//algoritma
//input(x)
//for j = 1 to 9 do
//input(bilangan)
//if bilangan = x then
//nX = nX + 1
//else
//nZero = nZero + 1
//endif
//endfor
//if nX >= nZero then
//output("Modus = ", x)
//else
//output("Modus = 0")
//endif
//endprogram