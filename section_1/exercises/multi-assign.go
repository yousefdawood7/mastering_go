package main

import "fmt"

func main() {
	var (
		lang    string
		version int
	)

	var (
		planet string  = "Mars"
		isTrue bool    = true
		temp   float64 = 19.5
	)

	lang, version = "Go", 2

	fmt.Println(lang, "version", version)

	fmt.Println("Air is good on " + planet)
	fmt.Println("it's ", isTrue)
	fmt.Println("it's ", temp)
}
