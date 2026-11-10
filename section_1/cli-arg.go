package main

import (
	"fmt"
	"os"
)

func main() {
	acceptedArg := os.Args[1:]

	fmt.Println(acceptedArg)
	fmt.Println("Temp Runned Path: " + os.Args[0])

	fmt.Println("Hello: " + acceptedArg[0])
}
