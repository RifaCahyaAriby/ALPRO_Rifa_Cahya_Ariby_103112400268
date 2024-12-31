package main 

import "fmt"

func main () {
	
	var i, y, bilangan, total int

	fmt.Scan(&y)
	for i = 0; i < 9; i++ {
		fmt.Scan(&bilangan)
		total += bilangan
	}
	if total >= y * 5{
		fmt.Println("Median bernilai", y)
	} else {
		fmt.Println("Median bernilai 0")
	}
}
//PSEUDOCODE
//program median
//kamus
//i, y, bilangan, total: integer
//algoritma
//input(y)
//for i = 0 to 9 do
//input(bilangan)
//total = total + bilangan
//endfor
//if total >= y * 5 then
//output("Median bernilai", y)
//else
//output("Median bernilai 0")
//endif
//endprogram