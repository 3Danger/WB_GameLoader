package handler

import (
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/httpserv/database"
	"log"
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
	Login() string
}

type Operator struct {
	sync.RWMutex
	accounts map[string]IAccount
	db       *database.DB
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

func NewOperator(db *database.DB) *Operator {
	return &Operator{
		accounts: make(map[string]IAccount),
		db:       db,
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

func (o *Operator) AddLoader(l *loader.Loader) {
	if ok := o.db.InsertLoader(l); ok != nil {
		log.Fatalln(ok)
	}
	o.Lock()
	o.accounts[l.Login()] = l
	o.Unlock()
}
func (o *Operator) AddCustomer(c *customer.Customer) {
	if ok := o.db.InsertCustomer(c); ok != nil {
		log.Fatalln(ok)
	}
	o.Lock()
	o.accounts[c.Login()] = c
	o.Unlock()
}
