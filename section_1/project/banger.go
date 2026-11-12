package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("You Gotta Pass at 1 arg")
		return
	}

	fmt.Println(strings.ToUpper(args[0]) + strings.Repeat("!", utf8.RuneCountInString(args[0])))
}
