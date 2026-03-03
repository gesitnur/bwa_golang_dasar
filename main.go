package main

import (
	"fmt"
	"pertama/calculation"
	"pertama/calculation_two"
)

func main() {
	fmt.Println("Hello, World!")

	result := calculation.Add(10, 20)
	fmt.Println(result)

	result_two := calculation_two.Multiply(10, 20)
	fmt.Println(result_two)
}
