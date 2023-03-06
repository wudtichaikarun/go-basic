package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
 total time = 101.3965ms
something
something
something
101.3965ms
*/
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	start := time.Now()

	go doSomeThing(wg)
	go doSomeThing(wg)
	go doSomeThing(wg)

	// ไม่ต้องเดา รอให้ทำงานเสร็จจริงๆ
	wg.Wait()

	fmt.Println(time.Since(start))
}

func doSomeThing(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something")
}
