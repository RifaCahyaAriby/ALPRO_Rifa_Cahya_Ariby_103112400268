package main
import "fmt"
func main () {
	var negara string
	var umur int
	
	fmt.Println("Syarat Pemilu")

	fmt.Print("Umur: ")
	fmt.Scan(&umur)
	fmt.Print("Kewarganegaraan (WNI/WNA): ")
	fmt.Scan(&negara)

	if umur >= 17 && negara == "WNI" {
		fmt.Println("Anda bisa mengikuti Pemilu")
	} else {
		fmt.Println("Anda tidak bisa mengikuti Pemilu!")
	}
	if umur < 17 {
		fmt.Println("Anda masih dibawah umur, minimal umur yaitu 17 tahun")
	}
	if negara != "WNI" {
		fmt.Println("Anda bukan WNI silahkan pindah negara")
	}

}