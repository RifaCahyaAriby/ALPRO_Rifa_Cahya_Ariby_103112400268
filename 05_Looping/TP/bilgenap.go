package main

import "fmt"

func main() {
    fmt.Println("Bilangan genap dari 1 hingga 50:")

    // Loop untuk menampilkan bilangan genap dari 1 hingga 50
    for i := 2; i <= 50; i += 2 {
        fmt.Print(i, " ")
    }
    fmt.Println() 
}
