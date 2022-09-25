package main

import "fmt"

func main() {
	d := 2
	double(&d)

	// 4
	fmt.Println(d)
}

/*
var s string
// <--- ประกาศ p เพื่อมาเก็บ address ของ string ตัวอื่น
var p * string

// <---- ref p to address ของ s
p = &s

// <---- assign value s ผ่านตัวแปร p
*p = "hello world"
*/

func double(d *int) {
	*d = *d * 2
}
