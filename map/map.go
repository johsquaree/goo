
package main

import "fmt"

func main() {

var mapExample map[string]int=map[string]int{"elma":493,"armut":372,"kavun":543}
for k,v:=range mapExample{
	fmt.Println(k,v)
}
/*MAPE KEY-VALUE EKLEME
mapExample["patates"]=781
for k,v:=range mapExample{
	fmt.Println(k,v)
}
*/
// KEY DAHA ÖNCEDEN KULLANIMIŞSA KEYİN DEĞERİ DEĞİŞİR ESKİ DEĞERİ GÖSTERİLMEZ

/* KEY SİLME
delete(mapExample,"kavun")
for k,v:=range mapExample{
	fmt.Println(k,v)
}
*/
// MAPTE İNDEX MANTIĞI YOKTUR KEYLER RASGELE GELİR.



}