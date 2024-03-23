package main

import (
	"fmt"

	"calculator.com/pkg/calculator"
)

func main() {
	fmt.Println("Running tax calculator!")
	calculator.RunCalculator("data2/prices.txt", "data/result-tax-")
}
