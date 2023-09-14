package main

import (
	"fmt"
	"math/rand"
	"time"
	
)
func main(){
	rand.Seed(time.Now().UnixNano())
	r:=rand.Intn(10)
	b:=rand.Intn(10)
	redGamer:=[]int{r,r,r,r,r,r,r,r,r,r}
	blueGamer:=[]int{b,b,b,b,b,b,b,b,b,b}
	
	
	/*KARTLARI ÖĞRENME*/
	fmt.Printf("%d\n",redGamer[3])
	fmt.Printf("%d\n",blueGamer[3])
	
	//oyunculların kartlarını görme
	//fmt.Printf("%v \n",blueGamer)
}

