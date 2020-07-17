package pointers

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is the error message of withdraw method
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin is a struct of a integer attribute
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is a object with a Bitcoin attribute
type Wallet struct {
	balance Bitcoin
}

// Deposit is method to add amounts to the wallet balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance is method to return the value in the balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw is method to remove amounts to the wallet balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
