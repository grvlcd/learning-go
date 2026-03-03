package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}

	prices := [4]float64{10.99, 9.99, 20.50, 12.99}

	productNames[2] = "A Carpet"

	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]

	fmt.Println(prices[2])
	fmt.Println(productNames)
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
