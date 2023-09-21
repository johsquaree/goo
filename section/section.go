package main

import (
    "fmt"
)

func main() {
    var sayi int
    fmt.Print("Pozitif bir tamsayi giriniz: ")
    fmt.Scan(&sayi)

    if sayi <= 0 {
        fmt.Println("Lütfen pozitif bir tamsayı giriniz.")
        return
    }

    fmt.Printf("%d sayisinin tam bölenleri: ", sayi)

    for i := 1; i <= sayi; i++ {
        if sayi%i == 0 {
            fmt.Printf("%d,", i)
        }
    }
    fmt.Println()
}
