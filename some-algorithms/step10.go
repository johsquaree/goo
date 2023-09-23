//Bir dizi verildiğinde, dizide hangi elemandan kaç tane olduğunu veren bir algoritma şeması hazırlayınız

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

	for i:=0;i<len(dizi);i++ {
	secilenIndex := dizi[i]
	a:=0
		for j:=0;j<len(dizi);j++ {
            if dizi[j]==secilenIndex{
				a++
			}
        }
		fmt.Println(secilenIndex,":",a)
    }
}