//Kullanıcı tarafından oluşturulan bir dizideki elemanları büyükten küçüğe doğru
//sıralayan algoritmayı ve akış şemasını hazırlayınız

package main

import (
    "fmt"
)

func main() {
fmt.Println("lütfen 5 indexlik diziniz için değerleri giriniz.")
var dizi[]int
	for i:=0;i<5;i++ {
		var a int
		fmt.Scan(&a)
		dizi=append(dizi,a)
	}
selectionSort(dizi)
fmt.Println(dizi)
}

func selectionSort(a []int) {
	n := len(a)
		for i := 0; i < n-1; i++ {
			maxIndex := i
			
			for j := i + 1; j < n; j++ {
				if a[j] > a[maxIndex] {
					maxIndex = j
				}
			}
		a[i], a[maxIndex] = a[maxIndex], a[i]
	}

}