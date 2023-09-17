package main
import "fmt"

func main() {
	var i int = 321
	j := &i
	
	fmt.Println("i address:", j) // i address: 0x14000124008
	fmt.Println("i value:", *j)  // 1 value: 2
  }