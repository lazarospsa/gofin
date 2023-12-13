package main

import (
	"fmt"
	gofin "github.com/lazarospsa/gofin"
)

func main() {
	fmt.Println(gofin.FutureValue(1000, 0.05, 10))
}