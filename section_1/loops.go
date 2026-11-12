package main

import "fmt"

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
}
