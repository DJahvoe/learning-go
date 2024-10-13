package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateEarnings(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.1f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.3f\n", ratio)

	writeResultToFile(fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", earningsBeforeTax, earningsAfterTax, ratio))
}

func writeResultToFile(resultString string) {
	os.WriteFile("result.txt", []byte(resultString), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("input can't be 0 or negative")
	}

	return userInput, nil
}

func calculateEarnings(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit
	return ebt, profit, ratio
}
