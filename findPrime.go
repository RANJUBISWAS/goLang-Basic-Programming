package main

import (
	"fmt"
)

func prime(num int) {
	if num < 2 {
		fmt.Println("No prime number.")
		return
	} else {
		i := 2
		for i < num+1 {
			p := true
			j := 2
			for j < i/2 {
				if i%j == 0 {
					p = false
				}
				j++
			}
			if p {
				fmt.Print(i, " ")
			}
			i++
		}
	}
}

func main() {
	fmt.Println("Enter number till which you want to see prime numbers:")
	var num int
	fmt.Scan(&num)
	prime(num)
}
