package loader

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/businesslogic/wallet"
	"errors"
)

type Builder struct {
	loader *Loader
}

func NewLoaderBuilder() *Builder {
	return &Builder{new(Loader)}
}

func (lb *Builder) AddAccount(acc *account.Account) *Builder {
	lb.loader.Account = acc
	return lb
}

func (lb *Builder) AddWallet(wallet *wallet.Wallet) *Builder {
	lb.loader.Wallet = wallet
	return lb
}

func (lb *Builder) AddTasks(task ...*task.Task) *Builder {
	for _, v := range task {
		lb.loader.tasks[v.Id] = v
	}
	return lb
}

func (lb *Builder) AddParams(maxWeightTrans, salary, fatigue float32, drunk bool) *Builder {
	lb.loader.maxWeightTrans = maxWeightTrans
	lb.loader.salary = salary
	lb.loader.fatigue = fatigue
	lb.loader.drunk = drunk
	return lb
}

func (lb *Builder) Build() (*Loader, error) {
	loader := lb.loader
	if loader.Account == nil {
		return nil, errors.New("account was not set")
	}
	if loader.Wallet == nil {
		return nil, errors.New("wallet was not set")
	}
	if loader.maxWeightTrans == 0. {
		return nil, errors.New("maxWeightTrans can't be zero")
	}
	if loader.salary == 0. {
		return nil, errors.New("salary can't be zero")
	}
	return loader, nil
}
