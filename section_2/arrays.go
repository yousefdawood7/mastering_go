package main

import (
	"fmt"
)

func main() {
	test := make([]string, 3, 5)

	// test = slices.Grow(test, 257)
	fmt.Println(cap(test))
	test = test[0:2:10]
	// fmt.Println(test[4])vvv
}
