package pointers

import (
	"errors"
	"fmt"
)

type Pikadollars int //my own, cool version of bitcoin

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Pikadollars
}

func (w *Wallet) Deposit(amount Pikadollars) { //pointer to the wallet
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Pikadollars) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Pikadollars {
	return w.balance
}

func (p Pikadollars) String() string {
	return fmt.Sprintf("%d PKD", p)
}
