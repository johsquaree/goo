// Girilen kilo ve boy bilgilerine göre Beden kitle/kütle indeksini hesaplayan 
// aynı zamanda bu sonuca göre kişinin ideal kilo durumunu ekrana
// yazdıran akış şeması

package main
import "fmt"

func main() {
	var w,h float32

	fmt.Print("VUCUT-KITLE ENDEXİNİZİ OGRENECESINIZ.\n")
    //metre ve cm farkı olabilir.ben kullanıcıan cm ile veri aldım.

	fmt.Print("boy katsayisini girin:\n(cm)")
	fmt.Scan(&h)
	fmt.Print("kilo katsayisini girin:\n(kilo) ")
	fmt.Scan(&w)

    kutleEndexi := w / (h * h)
    fmt.Println("Vücut Kitle İndeksi:", kutleEndexi)
    
    if kutleEndexi< 0.0018500 {
        fmt.Println("zayifsin yemek ye")
    } else if kutleEndexi< 0.0024999 {
        fmt.Println("normalsin ama abartma")
    } else if kutleEndexi< 0.0029999{
        fmt.Println("hafiften şişmansin dikkat et")
    } else if kutleEndexi< 0.0034999 {
        fmt.Println("şişmansin yemek yeme!")
    } else if kutleEndexi<0.0040000{
        fmt.Println("obezsin lütfen dur!")
    } else {
        fmt.Println("seni bu teste sokan şey ne?")
    }
}
