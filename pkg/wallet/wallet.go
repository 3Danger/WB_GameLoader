package wallet

import (
	. "GameLoaders/pkg/_interfaces"
	"errors"
)

type Wallet struct {
	money float32
}

func NewWallet(money float32) *Wallet { return &Wallet{money: money} }

func (w *Wallet) SendTo(money float32, wallet IWallet) error {
	if w.money < money {
		return errors.New("not enough money")
	}
	w.money -= money
	wallet.Receive(money)
	return nil
}

func (w Wallet) GetInfo() float32 {
	return w.money
}

func (w *Wallet) Receive(money float32) {
	w.money += money
}
