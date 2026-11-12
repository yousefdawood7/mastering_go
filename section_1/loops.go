package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("========")

Outerloop:
	for i := 0; i < 7; i++ {
		for j := 0; j < 13; j++ {
			if i%2 != 0 {
				continue Outerloop
			}
			fmt.Println(i, "*", j, "=", i*j)
		}
		fmt.Println("======")
	}

	items := []string{"Shoes", "Books", "Skate Boards"}

	for _, key := range items {
		fmt.Println("PRODUCT ===> ", key)
	}

	if lenItemZero := len(items[0]); lenItemZero == 5 {
		fmt.Println(strconv.Itoa(lenItemZero) + " Has a Exactly 5 chars")
	} else {
		fmt.Println(strconv.Itoa(lenItemZero) + " Has NOT a Exactly 5 chars")
	}
}
