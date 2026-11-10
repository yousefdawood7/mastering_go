package main

import (
	"fmt"
	"path"
)

func main() {
	dir, fileName := path.Split("/src/controller/app.controller.ts")

	fmt.Println(dir)
	fmt.Println(fileName)
}
