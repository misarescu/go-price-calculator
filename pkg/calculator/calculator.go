package calculator

import (
	"fmt"
	"sync"

	"calculator.com/pkg/filemanager"
)

func RunCalculator(filepath string) {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	// doneChans := make([]chan bool, len(taxRates))
	var wg sync.WaitGroup

	wg.Add(len(taxRates))

	for _, taxRate := range taxRates {
		// doneChans[idx] = make(chan bool)
		ioManager := filemanager.NewFileManager(filepath, fmt.Sprintf("data/result-tax-%.2f.json", taxRate))
		job := NewTaxIncludedPriceJob(*ioManager, taxRate)

		job.ProcessWG(&wg)
		// if err := job.Process(); err != nil {
		// 	fmt.Println("could not process job: ", err)
		// }
		// go job.ProcessChan(doneChans[idx])
	}
	// for _, d := range doneChans {
	// 	<-d
	// }

	wg.Wait()
}
