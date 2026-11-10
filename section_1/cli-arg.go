package main

import (
	"fmt"
	"os"
)

func main() {
	acceptedArg := os.Args[1:]

	fmt.Println(acceptedArg)

	fmt.Println("Hello: " + acceptedArg[0])
}
