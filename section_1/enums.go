package main

import "fmt"

func main() {
	type LogLevel int

	const (
		first LogLevel = iota
		second
		third
	)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}
