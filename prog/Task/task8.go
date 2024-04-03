package main

import "fmt"

func main() {
	var a int
	fmt.Print("Введите число от 1 до 359: ")
	fmt.Scan(&a)
	fmt.Println("It is", a*2/60, "hours", a*2%60, "minutes")
}
