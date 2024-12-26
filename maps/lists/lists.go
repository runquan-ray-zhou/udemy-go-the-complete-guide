package lists

import "fmt"

// type Product struct {
// 	id    string
// 	title string
// 	price float64
// }

// func main() {

// 	hobbies := [3]string{"reading manga", "watching films", "playing killer sudoku"}
// 	firstHobby := hobbies[0]
// 	secondHobby := hobbies[1]
// 	thirdHobby := hobbies[2]
// 	newList := [2]string{secondHobby, thirdHobby}

// 	sliceOne := hobbies[0:2]
// 	sliceTwo := hobbies[:2]

// 	reSlice := sliceOne[1:3]

// 	goals := []string{"understand Golang", "get certificate"}

// 	goals[1] = "complete by end of year"

// 	goals = append(goals, "build a project")

// 	products := []Product{
// 		{
// 			"first-product",
// 			"A First Product",
// 			12.99,
// 		},
// 		{
// 			"second-product",
// 			"A Second Product",
// 			129.99,
// 		},
// 	}

// 	newProduct := Product{
// 		"third-product",
// 		"A Third Product",
// 		15.99,
// 	}

// 	products = append(products, newProduct)

// 	fmt.Println(hobbies)
// 	fmt.Println(firstHobby)
// 	fmt.Println(newList)
// 	fmt.Println(sliceOne)
// 	fmt.Println(sliceTwo)
// 	fmt.Println(reSlice)
// 	fmt.Println(goals)
// 	fmt.Println(products)

// }

// Time to practice what you learned!

// [x]1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// [x]2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// [x]3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// [x]4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// [x]5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// [x]6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// [ ]7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

func main() {
	prices := []float64{10.99, 8.99} // this is a slice not an array. Go makes an array in the background
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	// updatedPrices := append(prices, 5.99)
	// fmt.Println(updatedPrices, prices)

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0} // creating an array
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	// featuredPrices := prices[1:3]
// 	// featuredPrices := prices[0:3]
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
