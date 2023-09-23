// Bir dizideki en küçük elemanı ekrana yazdıran algoritmayı ve akış şemasını hazırlayınız

package main

import (
    "fmt"
)
var a int

func main(){
	
	var dizi[]int = []int{-1,654654,-564,654,675,-3454325,432,2,645,32,21325,7563,123}
	a=dizi[0]
    for i:=0;i<len(dizi);i++{
        if dizi[i]<a{
           a=dizi[i]
        }
    }
	fmt.Println(a)
}