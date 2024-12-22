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

package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Revenue: ")
	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	expenses := getUserInput("Expenses: ")
	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	taxRate := getUserInput("Tax Rate: ")
	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	// fmt.Print(ebt, profit, ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancial(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return 
}