/*Kullanıcı tarafından girilen kilo ve boy bilgilerine göre vücut kitle indeksini hesaplayan
aynı zamanda bu sonuca göre kişinin ideal kilo durumunu ekrana yazdıran algoritmayı ve akış
şemasını hazırlayınız.
Not-1: Vücut kitle indeksi = Kilo / Boy * Boy
Not-2: Kilo => kg, Boy => m
Not-3: Vücut kitle indeksi < 18,5 ise Zayıf
 18,5 < Vücut kitle indeksi < 24,9 ise Normal
 25 < Vücut kitle indeksi < 29,9 ise Fazla kilolu
 30 < Vücut kitle indeksi < 34,9 ise I. derece obez
 35 < Vücut kitle indeksi < 39,9 ise II. derece obez
 Vücut kitle indeksi > 40 ise III. derece obez
*/
// Girilen kilo ve boy bilgilerine göre Beden kitle/kütle indeksini hesaplayan 
// aynı zamanda bu sonuca göre kişinin ideal kilo durumunu ekrana
// yazdıran akış şeması

package main
import "fmt"

func main() {
	var w,h float32

	fmt.Print("VUCUT-KITLE ENDEXİNİZİ OGRENECESINIZ.\n")
    //metre ve cm farkı olabilir.ben kullanıcıan m ile veri aldım.

	fmt.Print("boy katsayisini girin:\n(m)")
	fmt.Scan(&h)
	fmt.Print("kilo katsayisini girin:\n(kilo) ")
	fmt.Scan(&w)

    kutleEndexi := w / (h * h)
    fmt.Println("Vücut Kitle İndeksi:", kutleEndexi)
    
    if kutleEndexi< 18.5{
        fmt.Println("zayifsin yemek ye")
    } else if kutleEndexi< 24.999 {
        fmt.Println("normalsin ama abartma")
    } else if kutleEndexi< 29.999{
        fmt.Println("hafiften şişmansin dikkat et")
    } else if kutleEndexi< 34.999 {
        fmt.Println("şişmansin yemek yeme!")
    } else if kutleEndexi< 40.000{
        fmt.Println("obezsin lütfen dur!")
    } else {
        fmt.Println("seni bu teste sokan şey ne?")
    }
}

