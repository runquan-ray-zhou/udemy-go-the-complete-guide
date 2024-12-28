package main

import (
	"fmt"

	"github.com/runquan-ray-zhou/udemy-go-the-complete-guide/price_calculator/cmdmanager"
	"github.com/runquan-ray-zhou/udemy-go-the-complete-guide/price_calculator/prices"
)

type IOManager interface {
	Readlines() ([]string, error)
	WriteResult(data interface{}) error
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

}
