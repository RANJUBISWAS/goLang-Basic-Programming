package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter Number:")
	var num int
	fmt.Scan(&num)
	rev := 0
	temp := num
	for temp > 0 {
		rev = rev*10 + (temp % 10)
		temp /= 10
	}
	if num == rev {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a palindrome.")
	}
}
