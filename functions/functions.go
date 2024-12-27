package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)
	anotherSum := sumup(1, numbers...) //deconstruct the slice above

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// Variadic function, parameters are dynamic
func sumup(startingValue int, numbers ...int) int { // sumup can be called with a separate parameter values, as long as it is int
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
