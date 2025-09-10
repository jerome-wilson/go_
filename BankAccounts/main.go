package main

import (
	"bankaccounts/account"
	"fmt"
)
func main() {
	acc := account.BankAccount{Owner: "Jerome Wilson", Balance:943400}

	for {
		fmt.Printf("\n\n")
		fmt.Println("------------	WELCOME	  ------------")
		fmt.Println("Choose any option to proceed")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("q. Quit")
		fmt.Println("-------------------------------------")

		var choice string
		fmt.Print("\n")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			acc.AccountBalance()
		case "2":
			var amount float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scanln(&amount)
			err := acc.Deposit(amount)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "3":
			var amount float64
			fmt.Println("Enter the amount you want to withdraw: ")
			fmt.Scanln(&amount)
			err := acc.Withdraw(amount)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "q", "Q":
			fmt.Println("Thank you and have a nice day")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			
		}

	}
}