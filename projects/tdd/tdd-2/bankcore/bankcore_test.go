package bank

import "testing"

func TestAccount(t *testing.T)  {
	Account := Account {
		Customer : Customer {
			Name :"ariefs",
			Address: "Jakarta",
			Phone: "(021) 555 443",
		},
		Number: 2108,
		Balance: 0,
	}

	if Account.Number!= 2108 {
		t.Error("Account Number Error!")
	}
}

func TestDeposit(t *testing.T) {
	Account := Account {
		Customer : Customer {
			Name :"ariefs",
			Address: "Jakarta",
			Phone: "(021) 555 443",
		},
		Number: 2108,
		Balance: 0,
	}

	Account.Deposit(10)
	
	if Account.Balance != 10 {
		t.Error("Balance does not updated while deposit!")
	}
}

func TestWitdrawal(t *testing.T) {
	Account := Account {
		Customer : Customer {
			Name :"ariefs",
			Address: "Jakarta",
			Phone: "(021) 555 443",
		},
		Number: 2108,
		Balance: 0,
	}

	Account.Deposit(10)
	Account.Withdraw(10)
	
	if Account.Balance != 0 {
		t.Error("Withdraw failed!")
	}
}