/*
Ask for revenue, expense & tax rate
Calculate earnings before tax(EBT) and earnings after tax(profit)
Calculate ratio (EBT/profit)
Output EBT, profit and the ratio

Revenue is all income
EBT is profit after expense
Earning after taxes(Profit) is profit after taxes

To get EBT you subtract expenses from revenue
To get Taxes you multiply tax rate to EBT
To get profit you subtract taxes from EBT
To get ratio you divide profit into EBT
*/

// package main

// import "fmt"

// func main () {

// 	var revenue, expense, taxRate float64

// 	fmt.Print("What is your total revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Print("What is your total expense: ")
// 	fmt.Scan(&expense)

// 	fmt.Print("What is the current tax rate: ")
// 	fmt.Scan(&taxRate)

// 	earningsBeforeTax := revenue - expense
// 	taxes := earningsBeforeTax * (taxRate/100)
// 	profit := earningsBeforeTax - taxes
// 	ratio := earningsBeforeTax / profit

// 	fmt.Println(earningsBeforeTax)
// 	fmt.Println(profit)
// 	fmt.Println(ratio)

// }

// session 36 steps
// replace fmt.Print() and fmt.Scan() with a stand alone function
// that should output some text
// and then wait for the user to enter some text
// and then return that entered value.

//Goals
// 1) Validate user input
//    => Show error message & exit if invalid input is provided
//    - No negative numbers
//    - Not 0
// 2) Store calculated results into file

package main

import (
	"errors"
	"fmt"
	"os"
)

const profitFile = "profit.txt"

func storeResultsToFile(ebt, profit, ratio float64) {
	resultsText := fmt.Sprintf("ebt is %.1f, profit is %.1f, and ratio is %.3f", ebt, profit, ratio)
	os.WriteFile(profitFile, []byte(resultsText), 0644)
}

func main() {

	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		panic("Can't continue, sorry.")
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		panic("Can't continue, sorry.")
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		panic("Can't continue, sorry.")
	}

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	storeResultsToFile(ebt, profit, ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return userInput, errors.New("invalid input. must be greater than 0 and non negative")
	} else {
		return userInput, nil
	}
}

func calculateFinancial(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
