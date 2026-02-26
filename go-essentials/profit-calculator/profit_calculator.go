package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	printAndScan("Enter revenue: ", &revenue)

	printAndScan("Enter expenses: ", &expenses)

	printAndScan("Enter taxRate: ", &taxRate)

	earningsBeforeTax, earningsAfterTax, ratio := calculateEarningsAndRatio(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func printAndScan(label string, value *float64) {
	fmt.Print(label)
	fmt.Scan(value)
}

func calculateEarningsAndRatio(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt * (1 - taxRate/100)
	ratio := ebt / eat

	return ebt, eat, ratio
}
