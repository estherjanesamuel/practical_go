package bank

import "testing"

func TestAccount(t *testing.T)()  {
	account := Account{
		Customer : Customer{
			Name: "arief",
			Address: "Jakarta Raya, Indonesia",
			Phone: "(021) 555 0147",
		},
		Number: 1001,
		Balance: 0,
		}
		if account.Name == "" {
			t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T)()  {
	account := Account{
		Customer : Customer{
			Name: "arief",
			Address: "Jakarta Raya, Indonesia",
			Phone: "(021) 555 0147",
		},
		Number: 1001,
		Balance: 0,
		}

		account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalid(t *testing.T)()  {
	account := Account{
		Customer : Customer{
			Name: "arief",
			Address: "Jakarta Raya, Indonesia",
			Phone: "(021) 555 0147",
		},
		Number: 1001,
		Balance: 0,
		}

	if err := account.Deposit(-10); err == nil {
		t.Errorf("only positive numbers should be allowed to deposit %v", err)
	}
}

// Before we write the withdraw functionality, let's write the test for it.
func TestWithdraw(t *testing.T)  {
	account := Account{
		Customer : Customer{
			Name: "arief",
			Address: "Jakarta Raya, Indonesia",
			Phone: "(021) 555 0147",
		},
		Number: 1001,
		Balance: 0,
	}
	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer : Customer{
			Name: "arief",
			Address: "Jakarta Raya, Indonesia",
			Phone: "(021) 555 0147",
		},
		Number: 1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - arief - 100" {
		t.Error("statement doesn't have the proper format")
	}
}