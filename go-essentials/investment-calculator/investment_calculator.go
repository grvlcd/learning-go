package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	scanAndPrint("Enter investment amount: ", &investmentAmount)

	scanAndPrint("Enter investment year(s): ", &years)

	scanAndPrint("Enter investment return rate: ", &expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Investment Return %f \n", futureValue)
	fmt.Printf("Investment Real Value %f \n", futureRealValue)
}

func scanAndPrint(label string, value *float64) {
	fmt.Print(label)
	fmt.Scan(value)
}
