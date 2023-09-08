package main
import "fmt"

func tarih(x int,c int,v int){
	fmt.Println(x,"/",c,"/",v)
}
func saat(s int,d int){
	fmt.Println(s,":",d)
}
func main(){
	s:=18
	d:=35
	x:=23
	c:=4
	v:=3214
	tarih(x,c,v)
	saat(s,d)
}
