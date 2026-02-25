package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter taxRate: ")
	fmt.Scan(&taxRate)

	var earningsBeforeTax = revenue - expenses
	var earningsAfterTax = earningsBeforeTax * (1 - taxRate/100)
	var ratio = earningsBeforeTax / earningsAfterTax

	fmt.Printf("Earnings Before Tax: %f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %f\n", earningsAfterTax)
	fmt.Printf("Ratio: %f\n", ratio)
}
