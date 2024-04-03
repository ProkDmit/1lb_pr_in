package main

import "fmt"

func main() {
	var f, n, a int
	var t int = 10
	fmt.Print("Введите силу: ")
	fmt.Scan(&f)
	fmt.Print("Введите расстояние: ")
	fmt.Scan(&n)
	a = f * n
	fmt.Println("Работа: ", a)
	fmt.Println("Мощность: ", a/t)
}
