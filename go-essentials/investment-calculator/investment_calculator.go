package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	scanAndPrint("Enter investment amount: ", &investmentAmount)

	scanAndPrint("Enter investment year(s): ", &years)

	scanAndPrint("Enter investment return rate: ", &expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, years, expectedReturnRate)

	fmt.Printf("Investment Return %f \n", futureValue)
	fmt.Printf("Investment Real Value %f \n", futureRealValue)
}

func scanAndPrint(label string, value *float64) {
	fmt.Print(label)
	fmt.Scan(value)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}
