package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Print("DENKLEM SISTEMIMIZ\nax² + bx + c = 0\n--------------------------------\n")

	fmt.Print("a katsayisini girin: ")
	fmt.Scan(&a)
	fmt.Print("b katsayisini girin: ")
	fmt.Scan(&b)
	fmt.Print("c katsayisini girin: ")
	fmt.Scan(&c)

	if a == 0 {
		fmt.Println("a katsayisi sifir olamaz, ikinci dereceden bir denklem değil.")
		return
	}

	diskriminant := b*b - 4*a*c

	if diskriminant > 0 {
		//2kök var
		kok1 := (-b + math.Sqrt(diskriminant)) / (2 * a)
		kok2 := (-b - math.Sqrt(diskriminant)) / (2 * a)
		fmt.Printf("İki gerçel kök vardir: x1 = %.2f ve x2 = %.2f\n", kok1, kok2)
	} else if diskriminant == 0 {
		// İki kök eşittir ve reel sayıdır
		kok := -b / (2 * a)
		fmt.Printf("İki eşit ve gerçel kök vardir: x1 = x2 = %.2f\n", kok)
	} else {
		// Karmaşık kökler
		reelKisim := -b / (2 * a)
		imagKisim := math.Sqrt(-diskriminant) / (2 * a)
		fmt.Printf("Karmasik kökler vardir: x1 = %.2f + %.2fi ve x2 = %.2f - %.2fi\n", reelKisim, imagKisim, reelKisim, imagKisim)
	}
}
