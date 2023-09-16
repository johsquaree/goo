package main
import (
	"fmt"
	"math/rand"
	"time"	
)
func main(){
	rand.Seed(time.Now().UnixNano())
	redGamer:=[]int{rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10)}
	blueGamer:=[]int{rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10)}
	blueIndex := 0
	redIndex := 0

	fmt.Printf("Mavi Takim Ilk Indexinizi Seciniz:\n",)
	fmt.Scanf("%d",&blueIndex)
	time.Sleep(time.Second*3)
	fmt.Printf("Mavi takimin açilan ilk karti:%d\n",blueGamer[blueIndex])

	fmt.Printf("Kirmizi Takim Ilk Indexinizi Seciniz:\n")
	fmt.Scanf("%d",&redIndex)
	time.Sleep(time.Second*1)
	fmt.Printf("Kirmizi takimin acilan ilk karti:%d\n",redGamer[redIndex])
	
	time.Sleep(time.Second*1)
	//oyuncuların kartlarını görme
	fmt.Printf("%v \n",blueGamer)
	fmt.Printf("%v \n",redGamer)
}