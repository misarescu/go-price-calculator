package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"calculator.com/pkg/errors"
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
	// open file
	file, err := os.Open("data/prices.txt")

	if err != nil {
		file.Close()
		errors.ErrAndExit(err)
	}

	defer file.Close()

	// create scanner

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()
		floatNum, err := strconv.ParseFloat(line, len(line))
		if err != nil {
			errors.ErrAndExit(err)
		}

		job.InputPrices = append(job.InputPrices, floatNum)
	}

	if err := s.Err(); err != nil {
		errors.ErrAndExit(err)
	}
}
