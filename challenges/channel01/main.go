package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	exit := make(chan struct{})
	done := false

	runChannel := func(c chan int, value int) {
		go func() {
			for !done {
				c <- value
				milliSeconds := rand.Int() % 100
				time.Sleep(time.Duration(milliSeconds) * time.Millisecond)
			}
		}()
	}

	runChannel(ch1, 1)
	runChannel(ch2, 2)
	runChannel(ch3, 3)

	go func() {
		seconds := rand.Int() % 5
		time.Sleep(time.Duration(seconds) * time.Second)
		fmt.Println("send exit")
		exit <- struct{}{}
	}()

	for !done {
		select {
		case value := <-ch1:
			fmt.Println(value)
		case value := <-ch2:
			fmt.Println(value)
		case value := <-ch3:
			fmt.Println(value)
		case <-exit:
			fmt.Println("exit")
			done = true
		}
	}
	fmt.Println("done")
}
