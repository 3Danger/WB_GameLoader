package handler

import (
	"net/http"
	"sync"
	"time"
)

const (
	salt       = "#sdfjsdhdasd"
	signingKey = "#blabl$@#@#ablaqwe1231"
	tokenTTL   = time.Second * 10
)

type IAccount interface {
	GetUserName() string
	ToModel() interface{}
}

type Operator struct {
	sync.RWMutex
	accounts map[string]IAccount
}

func (o *Operator) GetRoute() *http.ServeMux {
	route := new(http.ServeMux)
	route.HandleFunc("/login", o.Login)
	route.HandleFunc("/register", o.Register)
	//route.HandleFunc("/tasks", o.Tasks)
	route.HandleFunc("/me", o.Me)
	//route.HandleFunc("/start", o.Start)
	return route
}

func NewOperator() *Operator {
	return &Operator{
		accounts: make(map[string]IAccount),
	}
}

func (o *Operator) GetUser(key string) IAccount {
	o.RLock()
	user := o.accounts[key]
	o.RUnlock()
	return user
}

func (o *Operator) HasLogin(login string) bool {
	o.RLock()
	_, ok := o.accounts[login]
	o.RUnlock()
	return ok
}

func (o *Operator) Add(user IAccount) {
	o.Lock()
	o.accounts[user.GetUserName()] = user
	o.Unlock()
}
