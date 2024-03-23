package calculator

import (
	"fmt"
	"sync"
	"time"

	"calculator.com/pkg/filemanager"
)

func RunCalculator(filepath, outputPrefix string) {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	errChan := make(chan error)
	var wg sync.WaitGroup

	wg.Add(len(taxRates))

	start := time.Now()

	for _, taxRate := range taxRates {
		ioManager := filemanager.NewFileManager(filepath, fmt.Sprintf("%s%.2f.json", outputPrefix, taxRate))
		job := NewTaxIncludedPriceJob(*ioManager, taxRate)

		go job.ProcessWG(&wg, errChan)
	}

	// print the errors if they come
	go func() {
		for {
			err := <-errChan
			fmt.Println("could not process job: ", err)
		}
	}()

	wg.Wait()

	fmt.Printf("it took %s\n", time.Since(start))
}
