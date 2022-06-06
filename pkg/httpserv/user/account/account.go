package account

import "sync"

type Account struct {
	mut        sync.RWMutex
	id         string
	name       string
	username   string
	password   string
	isCustomer bool
}

func NewAccount(id, name, username, password string, isCustomer bool) *Account {
	return &Account{
		id:         id,
		name:       name,
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
