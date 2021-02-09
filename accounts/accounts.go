package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner   string
	balance int
}

var NoMoney = errors.New("Can't withdraw you are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

//Withdraw x amount your account
func (a *Account) Withraw(amount int) error {
	if a.balance < amount {
		return NoMoney
	}
	a.balance -= amount
	return nil
}

//Balance return your account
func (a Account) Balance() int {
	return a.balance
}
