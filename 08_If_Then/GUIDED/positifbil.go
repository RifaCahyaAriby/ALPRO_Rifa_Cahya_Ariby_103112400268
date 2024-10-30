package main 

import "fmt"

func main() {
	var bil int 
	var teks string
	fmt.Scan(&bil)

	teks="bukan positif"
	if bil > 0 {
		teks="positif"
	}

	fmt.Println(teks)
	}
