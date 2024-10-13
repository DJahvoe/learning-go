package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Year of Investment: ")
	fmt.Scan(&years)

	fmt.Print("Expected Annual Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (inflation-adjusted): %.2f\n", futureRealValue)

	// fmt.Println("Future Value:", futureValue)
	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Println("Future Value (inflation-adjusted):", futureRealValue)
	// fmt.Printf("Future Value (inflation-adjusted): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// fmt.Printf(`Future Value: %.2f
	// Future Value (inflation-adjusted): %.2f
	// `, futureValue, futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv := fv / math.Pow((1+inflationRate/100), years)
	return fv, rfv
}
