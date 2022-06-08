package handler

import (
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
	customers map[string]IAccount
	loaders   map[string]IAccount
	tasks     []*task.Task
	db        *database.DB
}

func (o *Operator) GetRoute() *http.ServeMux {
	route := new(http.ServeMux)
	route.HandleFunc("/login", o.Login)
	route.HandleFunc("/register", o.Register)
	route.HandleFunc("/tasks", o.Tasks)
	route.HandleFunc("/me", o.Me)
	//route.HandleFunc("/start", o.Start)
	return route
}

func NewOperator(db *database.DB) *Operator {
	return &Operator{
		customers: make(map[string]IAccount),
		loaders:   make(map[string]IAccount),
		tasks:     make([]*task.Task, 0),
		db:        db,
	}
}

func (o *Operator) GetUser(key string) IAccount {
	o.RLock()
	user, ok := o.customers[key]
	if !ok {
		user, ok = o.loaders[key]
	}
	o.RUnlock()
	return user
}

func (o *Operator) HasLogin(login string) bool {
	o.RLock()
	_, ok := o.customers[login]
	if !ok {
		_, ok = o.loaders[login]
	}
	o.RUnlock()
	return ok
}

func (o *Operator) AddLoader(l *loader.Loader) {
	if ok := o.db.InsertLoader(l); ok != nil {
		log.Fatalln(ok)
	}
	o.Lock()
	o.loaders[l.Login()] = l
	o.Unlock()
}
func (o *Operator) AddCustomer(c *customer.Customer) {
	if ok := o.db.InsertCustomer(c); ok != nil {
		log.Fatalln(ok)
	}
	o.Lock()
	o.customers[c.Login()] = c
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
		customers[r].(*customer.Customer).AddTask(t)
	}
	o.Lock()
	o.tasks = append(o.tasks, tasks...)
	o.Unlock()
}
