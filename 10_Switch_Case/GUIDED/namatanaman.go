package main

import "fmt"

func main() {
    // masukan variabel dan tipedata
    var nama_tanaman string
    fmt.Scan(&nama_tanaman)
    switch nama_tanaman { // kondisi
    case "nepenthes", "drosera": // kondisi 1
        fmt.Println("Termasuk Tanaman Karnivora.") // cetak hasil
        fmt.Println("Asli Indonesia.")
    case "venus", "sarracenia": // kondisi 2
        fmt.Println("Termasuk Tanaman Karnivora.") // cetak hasil
        fmt.Println("Tidak Asli Indonesia.")
    default: // kondisi default
        fmt.Println("Tidak termasuk Tanaman Karnivora.") // cetak hasil
    }
}