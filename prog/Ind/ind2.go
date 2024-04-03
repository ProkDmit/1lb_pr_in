package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, h float64
	fmt.Print("1-ое основание: ")
	fmt.Scan(&a)
	fmt.Print("2-ое основание: ")
	fmt.Scan(&b)
	fmt.Print("Высота: ")
	fmt.Scan(&h)
	d = (b - a) / 2
	c = float64(math.Sqrt(h*h + d*d))
	fmt.Printf("%.2f", a+b+c*2)
}
