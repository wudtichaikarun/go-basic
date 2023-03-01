package main

import "fmt"

func main() {
	var s []int

	if s == nil {
		fmt.Println("it's nil")
		fmt.Println("++++++++++++++++++++++++++++++")
	}

	// []int is type of s
	// 4 is the length of the slice
	// 6 is the capacity of the slice
	s = make([]int, 4, 6)
	// value of s is [0 0 0 0]
	fmt.Println("emp:", s)
	// value of len(s) is 4
	fmt.Println("len:", len(s))
	// value of cap(s) is 6
	fmt.Println("cap(s)", cap(s))
	fmt.Println("++++++++++++++++++++++++++++++")

	s[0] = 1
	s[1] = 2
	// value of s is [1 2 0 0]
	fmt.Println("set:", s)
	fmt.Println("++++++++++++++++++++++++++++++")

	// append is a function that adds an element to the end of the slice
	s = append(s, 10)
	// after append the value of s is [1 2 0 0 10]
	fmt.Println("apd:", s)
	// value of len(s) is 5
	fmt.Println("len:", len(s))
	// value of cap(s) is 6
	fmt.Println("cap(s)", cap(s))
	fmt.Println("++++++++++++++++++++++++++++++")

	// a is array
	// [...] is a array literal
	a := [...]int{1, 3, 5, 7, 9}
	// value of a is [1 3 5 7 9]
	fmt.Println("a:", a)
	// mutate the slice
	s = a[:]
	// a: [1 3 5 7 9]
	// the pointer of s is the same as the pointer of a
	// s:0xc000016120 is same as a:0xc000016120

	// len:5 cap:5 s:0xc000016120 &a:0xc000016120
	fmt.Printf("len:%d cap:%d s:%p &a:%p\n", len(s), cap(s), s, &a)
	fmt.Println("++++++++++++++++++++++++++++++")

	s = append(s, 11, 13)
	// after append len and cap is more than maximum
	// go create new array and copy the old array to new array
	// now the pointer of s is changed to new array
	// s:0xc000014190 different from a:0xc000016120

	// len:7 cap:10 s:0xc000014190 &a:0xc000016120
	// after append the value of s is [1 2 0 0 10 11 13]
	fmt.Printf("len:%d cap:%d s:%p &a:%p\n", len(s), cap(s), s, &a)
	fmt.Println("++++++++++++++++++++++++++++++")
}
