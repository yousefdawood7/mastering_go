package main

import "fmt"

func main() {
	person := struct {
		id        int32
		firstName string
		lastName  string
	}{
		id:        1,
		firstName: "Yousef",
		lastName:  "Dawood",
	}

	fmt.Println(person)

	type Account struct {
		accountType string
		accountName string
	}

	yd7Account := Account{
		accountType: "Deposit",
		accountName: "YD7",
	}

	fmt.Println(yd7Account)
}
