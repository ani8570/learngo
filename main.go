package main

import (
	"fmt"

	"github.com/ani8570/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Lee")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withraw(60)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
