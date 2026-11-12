package main

import "fmt"

func main() {
	switch day := "Monday"; day {
	case "Monday":
		fmt.Println("OOOOOH YEAAA WE ARE IN MONDAY Y'ALL")
		fallthrough
	case "Friday":
		fmt.Println("Almost the weekend!")
	default:
		fmt.Println("Just a regular day.")
	}

	randomValue := 7
	switch {
	case randomValue == 7, randomValue == 10:
		fmt.Println("IT'S 7")
	}
}
