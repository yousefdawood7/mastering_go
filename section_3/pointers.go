package main

func stringPrinter(value *string) {
	println("PTR address =>", value)
	println("Value =>", *value)
}

func swapPTR(first *int, second *int) (*int, *int) {
	temp := *first // will be escaped to heap
	*first = *second
	second = &temp

	return first, second
}

func main() {
	a, b := 10, 20

	ptrStruct := &struct{ id int }{id: 1}
	ptrMap := &map[string]int{"id": 1}

	first, second := swapPTR(&a, &b)
	println(first, second)
	println(ptrStruct)
	println(ptrMap)
}
