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
/*	blueIndex := 0
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

	if redGamer[redIndex]==0 &&blueGamer[blueIndex]==0 {
		fmt.Printf("Iki Takimda Kazanmiştir\n")
	} else if redGamer[redIndex]==0{
		fmt.Printf("Kirmizi Takimda Kazanmiştir\n")
	} else if blueGamer[blueIndex]==0{
		fmt.Printf("Mavi Takimda Kazanmiştir\n")
	} else {
		fmt.Printf("Kartlar Yeniden Dağitiliyor\n")
	}
	
	yeniRedIndex:=redGamer[redIndex]
	yeniBlueIndex:=blueGamer[blueIndex]
	fmt.Printf("Mavi takimin açilan ilk karti:%d\n",blueGamer[yeniBlueIndex])
	fmt.Printf("Kirmizi takimin acilan ilk karti:%d\n",redGamer[yeniRedIndex])

	if yeniBlueIndex==0 &&yeniRedIndex==0 {
		fmt.Printf("Iki Takimda Kazanmiştir\n")
	} else if yeniBlueIndex==0{
		fmt.Printf("Mavi Takimda Kazanmiştir\n")
	} else if yeniRedIndex==0{
		fmt.Printf("Kirmizi Takimda Kazanmiştir\n")
	} else {
		fmt.Printf("Kartlar Yeniden Dağitiliyor\n")
	}

	yeni2RedIndex:=redGamer[yeniRedIndex]
	yeni2BlueIndex:=blueGamer[yeniBlueIndex]
	fmt.Printf("Mavi takimin açilan ilk karti:%d\n",blueGamer[yeni2BlueIndex])
	fmt.Printf("Kirmizi takimin acilan ilk karti:%d\n",redGamer[yeni2RedIndex])

	if yeni2BlueIndex==0 &&yeni2RedIndex==0 {
		fmt.Printf("Iki Takimda Kazanmiştir\n")
	} else if yeni2BlueIndex==0{
		fmt.Printf("Mavi Takimda Kazanmiştir\n")
	} else if yeni2RedIndex==0{
		fmt.Printf("Kirmizi Takimda Kazanmiştir\n")
	} else {
		fmt.Printf("Kartlar Yeniden Dağitiliyor\n")
	}

	yeni3RedIndex:=redGamer[yeni2RedIndex]
	yeni3BlueIndex:=blueGamer[yeni2BlueIndex]
	fmt.Printf("Mavi takimin açilan ilk karti:%d\n",blueGamer[yeni3BlueIndex])
	fmt.Printf("Kirmizi takimin acilan ilk karti:%d\n",redGamer[yeni3RedIndex])

	if yeni3BlueIndex==0 &&yeni3RedIndex==0 {
		fmt.Printf("Iki Takimda Kazanmiştir\n")
	} else if yeni3BlueIndex==0{
		fmt.Printf("Mavi Takimda Kazanmiştir\n")
	} else if yeni3RedIndex==0{
		fmt.Printf("Kirmizi Takimda Kazanmiştir\n")
	} else {
		fmt.Printf("Kartlar Yeniden Dağitiliyor\n")
	}
	*/
	b:=0
	fmt.Printf("Mavi Takim Ilk Indexinizi Seciniz:\n",)
	fmt.Scanf("%d",&b)
	time.Sleep(time.Second*1)
	for k:=1; k<11; k++{
		fmt.Printf("Mavi takimin acilan %d. karti:%d\n",k,blueGamer[b])
		if blueGamer[b]==0{
			fmt.Printf("Mavi Takimda Kazanmiştir\n")
			break
		}
		b=blueGamer[b]
	}

	time.Sleep(time.Second*5)
	r:=0
	fmt.Printf("Kirmizi Takim Ilk Indexinizi Seciniz:\n",)
	fmt.Scanf("%d",&r)
	time.Sleep(time.Second*1)

	for k:=1; k<11; k++{
		fmt.Printf("Kirmizi takimin acilan %d. karti:%d\n",k,redGamer[r])
		if redGamer[r]==0{
			fmt.Printf("Kirmizi Takimda Kazanmiştir\n")
			break
		}
		r=redGamer[r]
	}
	fmt.Printf("MAVI TAKIMIN KARTLARI:%v \n",blueGamer)
	fmt.Printf("KIRMIZI TAKIMIN KARTLARI:%v \n",redGamer)
}

