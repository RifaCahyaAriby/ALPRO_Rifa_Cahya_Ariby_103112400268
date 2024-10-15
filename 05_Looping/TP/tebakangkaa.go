package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Menginisialisasi generator angka acak
    rand.Seed(time.Now().UnixNano())

    // Memilih angka acak antara 1 hingga 100
    var target = rand.Intn(100) + 1
    var tebakan int
    var percobaan = 5 

    fmt.Println("Selamat datang di permainan tebak angka!")
    fmt.Println("Saya telah memilih sebuah angka antara 1 hingga 100.")
    fmt.Printf("Anda memiliki %d kesempatan untuk menebaknya.\n", percobaan)

    // memasukan 5 kesempatan
    for i := 1; i <= percobaan; i++ {
        fmt.Printf("Tebakan ke%d: Masukkan angka: ", i)
        fmt.Scan(&tebakan)

        // Mengecek apakah tebakan benar, terlalu kecil, atau terlalu besar
        if tebakan == target {
            fmt.Println("Selamat! Anda berhasil menebak angkanya dengan benar!")
            break
        } else if tebakan < target {
            fmt.Println("Tebakan Anda terlalu kecil.")
        } else {
            fmt.Println("Tebakan Anda terlalu besar.")
        }

        // Jika ini adalah kesempatan terakhir
        if i == percobaan {
            fmt.Printf("Maaf, Anda telah kehabisan kesempatan. Angka yang benar adalah %d.\n", target)
        }
    }

}
