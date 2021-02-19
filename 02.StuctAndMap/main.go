package main

import (
	"fmt"

	"github.com/ani8570/learngo/02.StuctAndMap/accounts"
	"github.com/ani8570/learngo/02.StuctAndMap/mydicts"
)

//account
func main() {
	account := accounts.NewAccount("Lee")

	account.Deposit(10)
	// err := account.Withdraw(60)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

	// }

	// // mydict
	// func mydict() {
	word := "first"

	dictionary := mydicts.Dictionary{word: "First word"}
	fmt.Println(dictionary[word])
	definition, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	if err := dictionary.Add("second", "Second word"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary)
	}

	if err := dictionary.Update(word, "Second word"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary)
	}

	dictionary.Delete(word)
	fmt.Println(dictionary)
}
