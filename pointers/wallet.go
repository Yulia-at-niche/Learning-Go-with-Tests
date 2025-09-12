package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin // Lowercase letters make something private to the package e.g 'b' in 'balance'.
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance() {
		// errors.New makes a new error with a custom message
		return errors.New("cannot complete transaction, insufficient funds")
	}
	w.balance -= amount
	return nil
}

/*
	A pointer is basically an address. Specifically, it's a variable that stores the memory address of where some data lives.

	A pointer points to the thing. It's an address to where the thing is in memory. This is similar to javascript reference types like objects and arrays.

	The big difference between Go and JS is that in Go, the default is to copy. You need to explicitly say "No, use the reference i.e. pointer".

	In JS it's the opposite. By default you access the reference and you need to explicitly copy and make a new reference (think of all the non mutating copying of state React and Redux)

	* Go: "Copy by default, opt into sharing with pointers"

	* JavaScript: "Share by default, opt into copying when needed"
*/
