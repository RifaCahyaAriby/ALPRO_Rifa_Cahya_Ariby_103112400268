package main 

import "fmt"

func main () {
	var kata, kunci string
 	kunci ="telkom" // kata kunci
	 
	for { 
		fmt.Print("Masukan kata: ") 
	 	fmt.Scan(&kata)
		fmt.Println("Anda memasukan kata :",kata)
		
		if kata == kunci { // jika "telkom" maka akan berhenti
		break
		}
	}
	fmt.Println("Program Selesai")
}
