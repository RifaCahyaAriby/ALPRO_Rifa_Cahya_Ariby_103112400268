package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan jumlah bintang: ")
    fmt.Scan(&n)

    // Mengubah menjadi segitiga bintang
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
