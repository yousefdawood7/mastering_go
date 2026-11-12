package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	firstName := "يوسف داود"

	fmt.Println(utf8.RuneCountInString(firstName))
}
