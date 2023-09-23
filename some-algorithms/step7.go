//:Bir öğrencinin almış olduğu vize notu ile final notunun ortalamasını hesaplayan
//algoritmayı ve akış şemasını hazırlayınız.(Vize notunun %30’unu final notunun %70’ini alarak
//hesaplama yapınız) (Notlar kullanıcı tarafından girilecektir)

package main

import "fmt"

func main() {
    var vize,ort,final float32
	fmt.Print("Vize notunu giriniz: ")
    fmt.Scan(&vize)
	fmt.Print("Final notunu giriniz: ")
    fmt.Scan(&final)

    ort = (vize * 0.3) + (final * 0.7)

    fmt.Println(ort)
}