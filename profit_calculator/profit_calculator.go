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

package main

import "fmt"

func main () {

	var revenue, expense, taxRate float64

	fmt.Print("What is your total revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("What is your total expense: ")
	fmt.Scan(&expense)

	fmt.Print("What is the current tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expense
	taxes := earningsBeforeTax * (taxRate/100)
	profit := earningsBeforeTax - taxes
	ratio := earningsBeforeTax / profit

	fmt.Println(earningsBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)

}