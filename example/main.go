package main

import (
	"fmt"
	"github.com/lazarospsa/gofin"
)

func main() {
	fmt.Println(gofin.futureValue(1000, 0.05, 10))
}