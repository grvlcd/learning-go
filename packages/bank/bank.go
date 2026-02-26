package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "./balance.txt"

func readBalanceToFile() (float64, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(content)
	data, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance.")
	}
	return data, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = readBalanceToFile()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Welcome to Go Bank!")
	for {
		var choice int
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your account balance is: %.2f\n", accountBalance)
		case 2:
			var depositAmount = 0.0
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your new account balance is: %.2f\n", accountBalance)
		case 3:
			var withdawAmount = 0.0
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdawAmount)

			if withdawAmount > accountBalance {
				fmt.Printf("Cannot withdraw this amount %.2f\n", withdawAmount)
				fmt.Printf("Your account balance is: %.2f\n", accountBalance)
				continue
			}

			accountBalance -= withdawAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your new account balance is: %.2f\n", accountBalance)
		case 4:
			fmt.Println("Thank you for choosing our bank!")
			return
		default:
			fmt.Println("Invalid choice!")
			continue
		}
	}
}
