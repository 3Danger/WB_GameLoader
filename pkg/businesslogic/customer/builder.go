package customer

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/businesslogic/wallet"
	"errors"
)

type Builder struct {
	customer *Customer
}

func NewCustomerBuilder() *Builder {
	return &Builder{new(Customer)}
}

func (b *Builder) AddAccount(acc *account.Account) *Builder {
	b.customer.Account = acc
	return b
}

func (b *Builder) AddWallet(wallet *wallet.Wallet) *Builder {
	b.customer.Wallet = wallet
	return b
}

func (b *Builder) AddTask(task ...*task.Task) *Builder {
	b.customer.tasks = append(b.customer.tasks, task...)
	return b
}

func (b *Builder) AddLoader(loader ...*loader.Loader) *Builder {
	b.customer.loaders = append(b.customer.loaders, loader...)
	return b
}

func (b *Builder) Build() (*Customer, error) {
	customer := b.customer
	if customer.Account == nil {
		return nil, errors.New("account was not set")
	}
	if customer.Wallet == nil {
		return nil, errors.New("wallet was not set")
	}
	return customer, nil
}
