package main

import "fmt"

func swapContents(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

func main() {
	list := []int{1, 2, 3}
	swapContents(list)
	fmt.Println(list)
}
