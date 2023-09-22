//1’den 100’e kadar olan sayıların toplamını ekrana yazdıran algoritmayı ve akış şemasını hazırlayınız

package main

import "fmt"

func main() {
    var toplam int
    for i := 1;  i<= 100; i++ {
        toplam += i
    }
    fmt.Println(toplam)
}