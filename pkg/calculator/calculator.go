package calculator

import (
	"fmt"

	"calculator.com/pkg/filemanager"
)

func RunCalculator(filepath string) {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChans[idx] = make(chan bool)
		ioManager := filemanager.NewFileManager(filepath, fmt.Sprintf("data/result-tax-%.2f.json", taxRate))
		job := NewTaxIncludedPriceJob(*ioManager, taxRate)
		// if err := job.Process(); err != nil {
		// 	fmt.Println("could not process job: ", err)
		// }
		go job.Process(doneChans[idx])
	}
	for _, d := range doneChans {
		<-d
	}
}
