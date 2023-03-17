package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan int)
	go fibonacci(ch)

	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println("ch value:", value, "at index", i)
	}

	fmt.Println("script took:", time.Since(start))
}

func fibonacci(ch chan int) {
	a, b := 0, 1

	time.Sleep(1 * time.Second)

	for {
		ch <- a
		a, b = b, a+b
	}
}
