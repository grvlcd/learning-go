package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := printAndScan("Enter revenue: ")

	if err != nil {
		panic("Revenue cannot be less than zero")
	}

	expenses, err := printAndScan("Enter expenses: ")

	if err != nil {
		panic("Expenses cannot be less than zero")
	}

	taxRate, err := printAndScan("Enter taxRate: ")

	if err != nil {
		panic("Tax rate cannot be less than zero")
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateEarningsAndRatio(revenue, expenses, taxRate)
	writeFromFile(earningsBeforeTax, earningsAfterTax, ratio)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func writeFromFile(ebt, eat, ratio float64) {
	textFormat := fmt.Sprintf("%.2f, %.2f, %.2f\n", ebt, eat, ratio)
	os.WriteFile("example.txt", []byte(textFormat), 0644)
}

func printAndScan(label string) (float64, error) {
	fmt.Print(label)

	value := 0.0

	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("Value cannot be less than zero")
	}

	return value, nil
}

func calculateEarningsAndRatio(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt * (1 - taxRate/100)
	ratio := ebt / eat

	return ebt, eat, ratio
}
