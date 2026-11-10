package main

import "fmt"

func main() {
	integerValue, floatValue := 12, 2.5

	integerValue = int(floatValue * float64(integerValue))

	fmt.Println(integerValue)
}
