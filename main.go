package main

import (
	"fmt"
	"io"
)

func main() {
	defer fmt.Println("defer#1")
	defer fmt.Println("defer#2")
	defer fmt.Println("defer#3")

	fmt.Println("end")
	/* result
	end
	defer#3
	defer#2
	defer#1
	*/

	dontForget(4)
	/**
	result
	8
	4
	*/

	tricky()
	/**
	result
	3
	3
	3
	*/
}

// ข้อควรระวังในการใช้
func dontForget(n int) {
	fmt.Println(n)

	// anonymous function
	defer func() {
		fmt.Println(n)
	}()

	n += n

}

// ข้อควรระวังในการใช้
func tricky() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func usecase(r io.ReadCloser) {
	defer r.Close()

	// b, err := ioutil.ReadAll(r)
	// ...
}
