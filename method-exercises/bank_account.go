package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the sum must be positive")
	}
	acc.balance += amount
	return nil
}

// Withdraw снимает средства
func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the sum must be positive")
	}
	if amount > acc.balance {
		return errors.New("not enough funds")
	}
	acc.balance -= amount
	return nil
}

func (acc BankAccount) CheckBalance() float64 {
	return acc.balance
}

func (acc BankAccount) DisplayInfo() {
	fmt.Printf("Owner: %s\n", acc.owner)
	fmt.Printf("Balance: %.2f\n", acc.balance)
}

func main() {
	account := NewBankAccount("Alexander Ivanov", 1000.0)
	account.DisplayInfo()

	err := account.Deposit(500.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The account has been replenished by 500.0")
		account.DisplayInfo()
	}

	err = account.Withdraw(200.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Removed 200.0")
		account.DisplayInfo()
	}
}
