package main

import (
	"fmt"
	"learngo/mydict"
)

// //account main
// func main() {
// 	account := accounts.NewAccount("Lee")

// 	account.Deposit(10)
// 	// err := account.Withdraw(60)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	fmt.Println(account.Balance(), account.Owner())
// 	fmt.Println(account)

// }

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	fmt.Println(dictionary["first"])
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
