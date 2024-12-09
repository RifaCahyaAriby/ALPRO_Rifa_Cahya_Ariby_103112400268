package main 

import "fmt"

func main () {
	var n int
	fmt.Print("Masukan batas : ")
	fmt.Scan(&n)

	for i := 2;  i <= n; i++ { // menampilkan bilangan prima
		prima := true // menentukan bilangan prima
		for j := 2; j*j <= i; j++ { 
			if i%j == 0 {
				prima = false // menentukan bilangan bukan prima
				break
			}
		}
		if prima {
			fmt.Println(i)
		}
	}
}