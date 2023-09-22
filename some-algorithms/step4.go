// 1’den 100’e kadar olan tek sayıları ekrana yazdıran algoritmayı ve akış şemasını hazırlayınız.

package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if i%2 != 0 {
            fmt.Print(i,",")
        }
    }
}