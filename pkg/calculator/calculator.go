package calculator

import (
	"fmt"

	"calculator.com/pkg/filemanager"
)

func RunCalculator() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		ioManager := filemanager.NewFileManager("data/prices.txt", fmt.Sprintf("data/result-tax-%.2f.json", taxRate))
		job := NewTaxIncludedPriceJob(*ioManager, taxRate)
		job.Process()
	}
}
