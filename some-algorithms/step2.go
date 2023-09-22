//Bir sayının tek mi çift mi olduğunu ekrana yazdıran algoritmayı ve akış şemasını hazırlayınız.
package main

import (
    "fmt"
)
func main() {
	var sayi int
	fmt.Print("sayi giriniz: ")
    fmt.Scanf("%d", &sayi)
    if sayi%2 == 0 {
        fmt.Println("cift")
    } else {
        fmt.Println("tek")
    }
}
