package calculator

import (
	"fmt"
	"sync"

	"calculator.com/pkg/conversions"
	"calculator.com/pkg/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64                 `json:"taxRate,omitempty"`
	InputPrices       []float64               `json:"netPrices,omitempty"`
	TaxIncludedPrices map[string]string       `json:"grossPrices,omitempty"`
	IOManager         filemanager.FileManager `json:"-"`
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) ProcessChan(done chan bool) {
	err := job.LoadData()

	if err != nil {
		// return err
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Printf("%v\n", result)
	job.TaxIncludedPrices = result

	job.IOManager.WriteJSON(job)
	done <- true
}

func (job *TaxIncludedPriceJob) ProcessWG(wg *sync.WaitGroup) {
	defer wg.Done()
	err := job.LoadData()

	if err != nil {
		// return err
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Printf("%v\n", result)
	job.TaxIncludedPrices = result

	job.IOManager.WriteJSON(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		// errors.ErrAndExit(err)
		return err
	}

	prices, err := conversions.StringsToFloats(lines)
	if err != nil {
		// errors.ErrAndExit(err)
		return err
	}

	job.InputPrices = prices
	return nil
}
