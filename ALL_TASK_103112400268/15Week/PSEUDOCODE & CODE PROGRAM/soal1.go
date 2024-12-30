package main

import "fmt"

func main (){
	var shimer, shine int 
	fmt.Scan(&shimer, &shine)

	if (shimer%2 == 0 && shine%2 != 0) || (shimer%2 != 0 && shine%2 == 0) { // menentukan ganjil genap atau sebaliknya
		fmt.Println("Berhasil")
	}

}
//PSEUDOCODE
//program ShimmerShine 
//kamus 
//tk1, tk2: integer 
//result: boolean  
//algoritma
//input(tk1, tk2) 
//hasil (tk1 + tk2) mod 2 == 0  
//if hasil then 
//output("Berhasil")  
//endif  
//endprogram
	

	 