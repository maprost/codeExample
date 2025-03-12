package main

import "fmt"

/*
	Write a GO Program to find factorial of a given number.
	Factorial of a number is the product of multiplication of a number n with every preceding number till it reaches 1. Factorial of 0 is 1.
*/

func main() {
	for n := 0; n <= 10; n++ {
		fmt.Printf("fac %d = %d | %d\n", n, facLoop(n), facRec(n))
	}
}

func facLoop(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func facRec(n int) int {
	if n == 0 {
		return 1
	}
	return n * facRec(n-1)
}
