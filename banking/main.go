package main

import (
	"fmt"

	"github.com/fernok/banking/accounts"
)

func main() {
	account := accounts.NewAccount("fernok")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(20) // error happens here
	// but you need to write code to handle that error
	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err)
		// calls Println and kills program
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
	account.ChangeOwner("donald")
	fmt.Println(account.Owner())
}
