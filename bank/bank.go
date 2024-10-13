package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	balanceText := string(data)

	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue the app")
		return
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println()
		fmt.Println("Action:")
		fmt.Println("1. Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Current Balance: ", accountBalance)
		case 2:
			fmt.Print("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		case 3:
			fmt.Print("Withdraw Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount.")
				continue
			}

			if accountBalance < withdrawAmount {
				fmt.Println("Balance insufficient.")
				continue
			}

			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
			// break
		}
	}

}
