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

func main () {
	const inflationRate = 2.5
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount) // pass pointer to scan "&", can not set constants

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)


	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	fmt.Println(futureValue) // Println make it easier to read
	fmt.Println((futureRealValue))
}