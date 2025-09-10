package account

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner string
	Balance float64
}

func (acc *BankAccount) Deposit (amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit value should be at least 1")
	}

	acc.Balance += amount
	fmt.Printf("Update successful! Your current balance is: %.2f\n", acc.Balance)
	return nil
}

func (acc *BankAccount) Withdraw(withdrawAmount float64) error {
	if withdrawAmount > acc.Balance {
		return errors.New("Insufficient Balance")
	}

	acc.Balance -= withdrawAmount
	fmt.Printf("Withdrawal successful! Your current balance is: %.2f\n", acc.Balance)
	return nil
}

func (acc BankAccount) AccountBalance() {
	fmt.Printf("Your account balance is: %.2f\n",acc.Balance)
}