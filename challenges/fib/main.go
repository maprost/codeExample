package main

import "fmt"

/*
	Write a Go program to find the nth Fibonacci number.

	To find the nth Fibonacci number, we have to add the previous 2 Fibonacci numbers as shown below.

	fib(0)=0
	fib(1)=1
	fib(2)=1+0 = 1
	fib(3)=1+1 = 2
	fib(4)=2+1 = 3
	:
	:
	fib(n)=fib(n-1)+fib(n-2)
*/

func main() {
	for n := 0; n <= 10; n++ {
		fmt.Printf("fac %d = %d\n", n, fibRec(n))
	}
}

func fibRec(n int) int {
	if n < 2 {
		return n
	}
	return fibRec(n-1) + fibRec(n-2)
}
