package main

import (
	"fmt"
	"time"
)

/**
ถ้า code is sequence
doSomeThing() 1 ครั้งทำงาน 100 Millisecond
doSomeThing() 1 ครั้งทำงาน 100 Millisecond
doSomeThing() 1 ครั้งทำงาน 100 Millisecond

sleep 120 Millisecond

รวมเท่ากับ 420 Millisecond ขึ้นไป

แต่ พอใช้ goroutine
รวมเท่ากับ 121.081875 Millisecond
*/

func main() {
	start := time.Now()

	go doSomeThing()
	go doSomeThing()
	go doSomeThing()

	time.Sleep(120 * time.Millisecond)

	fmt.Println(time.Since(start))
}

func doSomeThing() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something")
}
