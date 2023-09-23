//Bir dizinin ortalamasını bulan algoritmayı ve akış şemasını hazırlayınız

package main

import "fmt"

func main() {
	var ort,sayi float32
    var dizi[]float32=[]float32{-12432454,12,118,192,23,24,25,26,27,282,33,34,4564,56465,45645,454654,4546546,454,654,65,465,4,654,4564,}
	for i:=0;i<len(dizi);i++{
        sayi=dizi[i]+sayi
    }
	ort=sayi/float32(len(dizi))
	fmt.Println(ort)
}
