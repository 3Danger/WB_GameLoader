package handler

import (
	"GameLoaders/pkg/httpserv/user/customerAcc"
	"GameLoaders/pkg/httpserv/user/loaderAcc"
)

type Operator struct {
	customer map[string]*customerAcc.User
	loader   map[string]*loaderAcc.User
}

func NewOperator() *Operator {
	return &Operator{
		customer: make(map[string]*customerAcc.User),
		loader:   make(map[string]*loaderAcc.User),
	}
}

func (o *Operator) GetCustomer(key string) *customerAcc.User {
	return o.customer[key]
}

func (o *Operator) GetLoader(key string) *loaderAcc.User {
	return o.loader[key]
}

type iAccount interface {
	IsCustomer() bool
	Id() string
}

func (o *Operator) Add(user iAccount) {
	if user.IsCustomer() {
		//TODO key!!!
		o.customer[user.Id()] = user.(*customerAcc.User)
	} else {
		//TODO key!!!
		o.loader[user.Id()] = user.(*loaderAcc.User)
	}
}
