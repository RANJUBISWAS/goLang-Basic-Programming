package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Print("Enter a Number:")
	fmt.Scanln(&a)
	sum := 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	fmt.Println("Sum:", sum)
}
