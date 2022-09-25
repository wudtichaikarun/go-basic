package main

import (
	"fmt"

	"github.com/wudtichaikarun/hello/book"
)

func main() {
	b := book.New()

	fmt.Printf("%T %v\n", b, b)

	b.Name = "Romantic"
	fmt.Printf("%T %v\n", b, b)
}
