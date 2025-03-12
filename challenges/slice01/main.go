package main

import "fmt"

func main() {
	var x []int      // x=[], len=0, cap=0
	x = append(x, 1) // x=[1], len=1, cap=1
	x = append(x, 2) // x=[1, 2], len=2, cap=2
	x = append(x, 3) // x=[1, 2, 3], len=3, cap=4
	y := x           // y=[1, 2, 3], len=3, cap=4
	x = append(x, 4) // x=[1, 2, 3, 4], len=4, cap=4
	y = append(y, 5) // x=[1, 2, 3, 5], len=4, cap=4
	x[0] = 0         // x=[0, 2, 3, 5], len=4, cap=4

	fmt.Println(x) // x=[0, 2, 3, 5], len=4, cap=4
	fmt.Println(y) // y=[0, 2, 3, 5], len=4, cap=4
}
