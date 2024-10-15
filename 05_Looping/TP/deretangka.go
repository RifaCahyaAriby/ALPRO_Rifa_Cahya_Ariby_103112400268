package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&n)

    // Menghitung jumlah deret angka dari 1 hingga n
    var total int = 0
    for i := 1; i <= n; i++ {
        total += i
		fmt.Println(i)
    }
    fmt.Printf("Jumlah deret angka dari 1 hingga %d adalah %d\n", n, total)
}
