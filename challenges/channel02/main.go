package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(len(array))
	go func() {
		for _, value := range array {
			ch <- &value
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value)
			wg.Done()
		}
	}()

	// New goroutine is run.
	//go func() {
	//	var a int
	//	for { // will run forever -> programm can not stop
	//		a++
	//	}
	//}()

	wg.Wait()
}
