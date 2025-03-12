package main

import "fmt"

func main() {
	var x int
	arr := [3]int{3, 5, 2} // remove the 3, now it's a slice and it works
	x, arr = arr[0], arr[1:]
	fmt.Println(x, arr)
}
