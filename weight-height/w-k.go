// Girilen kilo ve boy bilgilerine göre Beden kitle/kütle indeksini hesaplayan 
// aynı zamanda bu sonuca göre kişinin ideal kilo durumunu ekrana
// yazdıran akış şeması

package main
import "fmt"

func main() {
    var input float64
    var w float64


    fmt.Println("Lütfen boy ve kilonuzu sirasiyla giriniz...")
    fmt.Println("Boy Bilgisi:")
    fmt.Scanf(&h)
    fmt.Println("Kilo Bilgisi:")
    fmt.Scanf(&w)

    kutleEndexi := w / (h * h)
    fmt.Println("Vücut Kitle İndeksi:", kutleEndexi)
}
