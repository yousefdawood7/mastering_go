package main

func stringPrinter(value *string) {
	println("PTR address =>", value)
	println("Value =>", *value)
}

func main() {
	fullName := "Yousef Dawood"

	stringPrinter(&fullName)
}
