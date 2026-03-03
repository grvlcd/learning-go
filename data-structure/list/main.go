package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)
	prices[1] = 9.99
	fmt.Println(prices)
	prices = append(prices, 30.99)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
//
// 	prices := [4]float64{10.99, 9.99, 20.50, 12.99}
//
// 	productNames[2] = "A Carpet"
//
// 	featuredPrices := prices[1:]
// 	highlightedPrices := featuredPrices[:1]
//
// 	fmt.Println(prices[2])
// 	fmt.Println(productNames)
// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
