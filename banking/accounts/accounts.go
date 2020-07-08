package accounts

import (
	"errors"
	"fmt"
)

// private struct - starts with lowercase
// the keys inside should also be uppercase to make it public
type Account struct {
	owner   string
	balance int
}

// error name should be errOOOO
var errNoMoney = errors.New("Can't Withdraw!")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Method --> func (receiver) functionName(attributes) returnValue {}
// the receiver should start with the lowercased first letter of the struct name
// add * to not make a copy but bring the actual struct
// in this case, the struct name is Account, so the receiver name should start with 'a'

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Error handling in Go

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
		// return errors.New("Can't withdraw: you are poor")
	}
	a.balance -= amount
	return nil
}

// ChangeOwner uppdates the owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner returns the owner of the account
func (a Account) Owner() string {
	return a.owner
}

// when you print a struct, the compiler returns &{ ~~~~ }
// this is a struct formatted in string
// and the compiler calls the String() method automatically
// so you can manually specify what the compiler will do when you print a struct
// by modifying the String() method

// String is a method that Go calls internally
func (a Account) String() string {
	// return "whatever you want"
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
