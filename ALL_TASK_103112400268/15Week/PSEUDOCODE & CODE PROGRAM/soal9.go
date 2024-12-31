package main 

import "fmt"

func main() {
	var p, r, hari int 

	fmt.Scan(&p, &r)
	hari = p / r 
	if p % r > 0 { 
		hari++ 
	}
	fmt.Println(hari, "hari")
}

//PSEUDOCODE
//program novel
//kamus
//p, r, hari: integer
//algoritma
//input(p, r)
//hari = p / r
//if p mod r > 0 then
//hari++
//endif
//output(hari, "hari")
//endprogram