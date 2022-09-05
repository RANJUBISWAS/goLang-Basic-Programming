package main

import (
	"fmt"
)

func factorial(num int) int {
	if num < 0 {
		return 0
	} else if num == 0 || num == 1 {
		return 1
	} else {
		return factorial(num-1) * num
	}
}
func main() {
	fmt.Print("Enter a number:")
	var num int
	fmt.Scan(&num)
	fact := factorial(num)
	fmt.Println("Factorial is:", fact)
}
