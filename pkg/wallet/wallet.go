package wallet

import (
	"errors"
)

type IWallet interface {
	SendTo(int, IWallet) error
	Get() int
	Receive(int)
}

type Wallet struct {
	money int
}

func NewWallet(money int) *Wallet { return &Wallet{money: money} }

func (w *Wallet) SendTo(money int, wallet IWallet) error {
	if w.money < money {
		return errors.New("insufficient funds")
	}
	w.money -= money
	wallet.Receive(money)
	return nil
}

func (w Wallet) Get() int {
	return w.money
}

func (w *Wallet) Receive(money int) {
	w.money += money
}
