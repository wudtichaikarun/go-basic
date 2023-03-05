package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello , Romantic")

	// เราสามารถใส่ short statement ก่อน condition
	// n, err := strconv.Atoi("5s")
	// if err != nil {
	// 	fmt.Println("NaN:", err)
	// } else {
	// 	fmt.Println(n)
	// }
	if n, err := strconv.Atoi("5s"); err != nil {
		fmt.Println("NaN:", err)
	} else {
		fmt.Println(n)
	}

}

func showName(name string) {
	if name != "" {
		fmt.Println("Hello $s\n", name)
	} else {
		fmt.Println("Hello friend")
	}
}
