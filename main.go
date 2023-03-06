package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	doSomeThing()
	time.Sleep(time.Second)
	fmt.Println("Hello , Romantic")
}

func doSomeThing() {
	fmt.Println("do some thing")
}
