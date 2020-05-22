package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is a global variable
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Stringer interface
type Stringer interface {
	String() string
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit allows to deposit amount to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	(*w).balance += amount
	// we can just write w.balance
	//struct pointers are automatically dereferenced
}

// Withdraw allows to withdraw amount from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance allows to get value from wallet
func (w *Wallet) Balance() Bitcoin {
	return (*w).balance
}
