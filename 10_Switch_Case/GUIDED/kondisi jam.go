package main

import "fmt"

func main() {

    // masukan variabel dan tipedata
    var jam12, jam24 int  
    var label string
    fmt.Scan(&jam24) // scan inputan jam
    switch { // kondisi
    case jam24 == 0: // kondisi 1
        jam12 = 12
        label = "AM" //
    case jam24 < 12: // kondisi 2
        jam12 = jam24
        label = "AM"
    case jam24 == 12: // kondisi 3
        jam12 = 12
        label = "PM"
    case jam24 > 12 && jam24 < 23: // kondisi 4
        jam12 = jam24 - 12
        label = "PM"
    default:    // kondisi default
        fmt.Print("Tidak termasuk jam")
    } 

    fmt.Println(jam12, label) // cetak hasil
}