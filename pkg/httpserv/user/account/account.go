package account

import "sync"

type IAccount interface {
	IsCustomer() bool
	Id() string
	Login() string
	GetAccount() IAccount
}

type Account struct {
	mut        sync.RWMutex
	id         string
	name       string
	username   string
	password   string
	isCustomer bool
}

func NewAccount(id, username, password string, isCustomer bool) *Account {
	return &Account{
		id:         id,
		username:   username,
		password:   password,
		isCustomer: isCustomer,
	}
}

func (a *Account) IsCustomer() bool {
	a.mut.RLock()
	defer a.mut.RUnlock()
	return a.isCustomer
}

func (a *Account) Id() string {
	a.mut.RLock()
	defer a.mut.RUnlock()
	return a.id
}

func (a *Account) Name() string { return a.name }

func (a *Account) Login() string {
	return a.username
}
func (a *Account) GetAccount() IAccount {
	return a
}
