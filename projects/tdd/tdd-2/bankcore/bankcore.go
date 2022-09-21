package bank

import "fmt"

type Account struct {
	Customer
	Number int
	Balance float64
}

type Customer struct {
	Name, Address, Phone string
}

func (a *Account)Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("an error occured While deposit %v to your account. Deposit must be greater than zero", amount)
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("an error occured While withdraw %v your account. withdraw must be greater than zero", amount)
	}

	if a.Balance < amount {
		return fmt.Errorf("an error occure while withdrawal %v your account. not enough balance", amount)
	}

	a.Balance -= amount
	return nil
}