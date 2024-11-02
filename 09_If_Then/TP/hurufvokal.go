package main
import "fmt"
func main() {
	var huruf string
	fmt.Print("Cek huruf: ")
	fmt.Scan(&huruf)

	if huruf == "A" || huruf == "a" ||
	huruf == "I" || huruf == "I" ||
	huruf == "U" || huruf == "u" ||
	huruf == "E" || huruf == "e" ||
	huruf == "O" || huruf == "o" {
	fmt.Println("Huruf di atas adalah Huruf Vokal")
} else {
	fmt.Println("Huruf diatas adalah Huruf Konsonan")
}
}