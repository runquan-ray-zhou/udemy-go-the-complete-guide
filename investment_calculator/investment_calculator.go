// package main //package instructions, every go code file.
// // use main because it is a special package name, that tells go this package will be the main entry point for the application we are building

// import "fmt" // importing fmt package, part of go standard library itself

// func main() {
// 	fmt.Print("Hello World")// calling a function, Print is a function/executing command, use Print feature
// }
// // string must be double quotes or back ticks

package main

import (
	"fmt" // get input from the terminal
	"math"
)

const inflationRate = 2.5

func main() {
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	// fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount) // pass pointer to scan "&", can not set constants

	outputText("Expected Return Rate: ")
	// fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	// fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	// outputs information
	// fmt.Println("Future Value:", futureValue) // Println make it easier to read
	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

// Better to state what you are returning
// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64 ) (float64, float64){
// 	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
// 	rfv := fv / math.Pow(1 + inflationRate/100, years)
// 	return fv, rfv
// }

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}
