package calculator

import (
	"fmt"

	"calculator.com/pkg/conversions"
	"calculator.com/pkg/errors"
	"calculator.com/pkg/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Printf("%v\n", result)
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("data/prices.txt")
	if err != nil {
		errors.ErrAndExit(err)
	}

	prices, err := conversions.StringsToFloats(lines)
	if err != nil {
		errors.ErrAndExit(err)
	}

	job.InputPrices = prices
}
