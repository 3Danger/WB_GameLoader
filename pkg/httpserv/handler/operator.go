package handler

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/httpserv/database"
	"log"
	"math/rand"
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
	Id() int
	Login() string
	ToModel() interface{}
	Tasks() []*task.Task
}

type Operator struct {
	sync.RWMutex
	//key = login
	accounts map[string]*account.Account
	//key = ID
	customers map[int]IAccount
	//key = ID
	loaders map[int]IAccount
	tasks   []*task.Task
	db      *database.DB
}

func (o *Operator) GetRoute() *http.ServeMux {
	route := new(http.ServeMux)
	route.HandleFunc("/login", o.Login)
	route.HandleFunc("/register", o.Register)
	route.HandleFunc("/tasks", o.Tasks)
	route.HandleFunc("/me", o.Me)
	route.HandleFunc("/start", o.Start)
	return route
}

func NewOperator(db *database.DB) *Operator {
	accounts := make(map[string]*account.Account)
	customers := make(map[int]IAccount)
	loaders := make(map[int]IAccount)
	aviableTasks := db.LoadTasks(0)

	for _, v := range db.LoadCustomers() {
		customers[v.Account.Id()] = v
		accounts[v.Account.Login()] = v.Account
	}
	for _, v := range db.LoadLoaders() {
		loaders[v.Account.Id()] = v
		accounts[v.Account.Login()] = v.Account
	}
	for _, v := range loaders {
		if c, ok := customers[v.(*loader.Loader).CustomerAccountId()]; ok {
			c.(*customer.Customer).HireLoader(v.(*loader.Loader))
		}
	}
	return &Operator{
		accounts:  accounts,
		customers: customers,
		loaders:   loaders,
		tasks:     aviableTasks,
		db:        db,
	}
}

func (o *Operator) GetUser(login string) IAccount {
	o.RLock()
	defer o.RUnlock()
	user, ok := o.accounts[login]
	if ok {
		if cstmr, ok := o.customers[user.Id()]; ok {
			return cstmr
		} else {
			return o.loaders[user.Id()]
		}
	}
	return nil
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
	o.accounts[l.Account.Login()] = l.Account
	o.loaders[l.Account.Id()] = l
	o.Unlock()
}
func (o *Operator) AddCustomer(c *customer.Customer) {
	if ok := o.db.InsertCustomer(c); ok != nil {
		log.Fatalln(ok)
	}
	o.Lock()
	o.accounts[c.Account.Login()] = c.Account
	o.customers[c.Account.Id()] = c
	o.Unlock()
}

func (o *Operator) AddTasks(tasks ...*task.Task) {
	var r int
	var customers []IAccount
	for _, v := range o.customers {
		customers = append(customers, v)
	}
	for _, t := range tasks {
		r = rand.Int() % len(o.customers)
		if ok := o.db.InsertTask(t, customers[r].(*customer.Customer).Account.Id()); ok != nil {
			log.Fatalln(ok)
		}
		customers[r].(*customer.Customer).AddTasks(t)
	}
	o.Lock()
	o.tasks = append(o.tasks, tasks...)
	o.Unlock()
}
