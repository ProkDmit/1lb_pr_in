package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("Введите значение a: ")
	fmt.Scan(&a)
	fmt.Print("Введите значение b: ")
	fmt.Scan(&b)
	s := 2 * math.Pi * a * b
	v := (4 / 3) * math.Pi * a * b * b
	fmt.Printf("Площадь = %1.2f \nОбъем = %1.2f", s, v)
}
