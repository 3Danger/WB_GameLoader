package wallet

import (
	. "GameLoaders/pkg/businesslogic/interfaces"
	"errors"
	"sync"
)

type Wallet struct {
	mut   *sync.RWMutex
	money float32
}

func NewWallet(money float32) *Wallet {
	return &Wallet{
		mut:   new(sync.RWMutex),
		money: money,
	}
}

func (w *Wallet) SendTo(money float32, wallet IWallet) error {
	w.mut.RLock()
	if w.money < money {
		w.mut.RUnlock()
		return errors.New("not enough money")
	}
	w.mut.RUnlock()
	w.mut.Lock()
	w.money -= money
	w.mut.Unlock()
	wallet.Receive(money)
	return nil
}

func (w Wallet) GetInfo() float32 {
	return w.money
}

func (w *Wallet) Receive(money float32) {
	w.mut.Lock()
	w.money += money
	w.mut.Unlock()
}
