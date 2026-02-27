package main

import (
	"example.com/bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const fileName = "./balance.txt"

func main() {
	var accountBalance, err = fileops.ReadToFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("Welcome to Go Bank! %v \n", randomdata.SillyName())
	for {
		var choice int
		presentOptions()

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
			fileops.WriteToFile(accountBalance, fileName)
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
			fileops.WriteToFile(accountBalance, fileName)
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
