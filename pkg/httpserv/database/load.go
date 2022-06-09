package database

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/businesslogic/wallet"
	"log"
	"strconv"
)

func (d *DB) LoadCustomers() (customers []*customer.Customer) {
	var (
		accountId  int
		login      string
		password   string
		customerId int
		money      float32
		cb         *customer.Builder
	)
	query := "SELECT account.id, account.login, account.password, customer.id, customer.money " +
		"FROM account INNER JOIN customer ON account.id = customer.account_id"
	rows, ok := d.connect.Query(query)
	if ok != nil {
		log.Fatalln(ok)
	}
	for rows.Next() {
		if ok = rows.Scan(&accountId, &login, &password, &customerId, &money); ok != nil {
			rows.Close()
			log.Fatalln(ok)
		}
		cb = customer.NewCustomerBuilder()
		cb.AddWallet(wallet.NewWallet(money))
		cb.AddAccount(account.NewAccount(accountId, login, password))
		cb.SetId(customerId)
		customers = append(customers, cb.Customer())
	}
	rows.Close()
	for _, c := range customers {
		c.AddTasks(d.LoadTasks(c.Account.Id())...)
	}

	return customers
}

func (d *DB) loadAccount(id int) *account.Account {
	var login, password string
	query := "SELECT login, password FROM account WHERE id = $1 "
	rows, ok := d.connect.Query(query, id)
	if ok != nil {
		log.Fatalln(ok)
	}
	defer rows.Close()
	if rows.Next() {
		if ok = rows.Scan(&login, &password); ok != nil {
			log.Fatalln(ok)
		}
		return account.NewAccount(id, login, password)
	}
	return nil
}

func (d *DB) LoadLoaders() (loaders []*loader.Loader) {
	var (
		accountId int
		login     string
		password  string

		id, customerId          int
		money, salary           float32
		maxWeightTrans, fatigue float32
		drunk                   bool

		lb *loader.Builder
	)

	//SELECT account.id, account.login,
	// loader.id, loader.money, loader.salary,
	// loader.max_weight_trans, loader.fatigue,
	// loader.drunk, FROM loader
	// INNER JOIN account ON loader.account_id = account.id

	query := "SELECT account.id, account.login, account.password, " +
		"loader.customer_id, loader.id, loader.money, loader.salary, loader.max_weight_trans, loader.fatigue, loader.drunk " +
		"FROM loader INNER JOIN account ON loader.account_id = account.id"
	rows, ok := d.connect.Query(query)
	if ok != nil {
		log.Fatalln(ok)
	}

	for rows.Next() {
		ok = rows.Scan(&accountId, &login, &password, &customerId, &id, &money, &salary, &maxWeightTrans, &fatigue, &drunk)
		if ok != nil {
			rows.Close()
			log.Fatalln(ok)
		}
		lb = loader.NewLoaderBuilder()
		lb.AddWallet(wallet.NewWallet(money))
		lb.AddAccount(account.NewAccount(accountId, login, password))
		lb.AddParams(id, customerId, maxWeightTrans, salary, fatigue, drunk)
		loaders = append(loaders, lb.Loader())
	}
	rows.Close()
	for _, v := range loaders {
		if v.Account.Id() > 0 {
			rows, ok = d.connect.Query("SELECT account.id FROM customer INNER JOIN account ON customer.account_id = account.id WHERE customer.id = $1", v.CustomerId())
			if ok != nil {
				log.Fatalln(ok)
			}
			if rows.Next() {
				customerAccountId := 0
				if ok = rows.Scan(&customerAccountId); ok != nil {
					log.Fatalln(ok)
				}
				v.SetCustomerAccountId(customerAccountId)
			}
			rows.Close()
			v.AddTask(d.LoadTasks(v.Account.Id())...)
		}
	}
	return loaders
}

func (d *DB) LoadTasks(accountId int) (tasks []*task.Task) {
	var (
		id     int
		name   string
		weight float32
	)
	tasks = make([]*task.Task, 0)
	query := "SELECT id, name, weight FROM tasks WHERE account_id = $1 "
	if rows, ok := d.connect.Query(query, accountId); ok != nil {
		log.Fatalln(ok)
	} else {
		defer rows.Close()
		for rows.Next() {
			if ok = rows.Scan(&id, &name, &weight); ok != nil {
				log.Fatalln(ok)
			}
			tasks = append(tasks, &task.Task{
				Id:     strconv.Itoa(id),
				Name:   name,
				Weight: weight,
			})
		}
	}
	return tasks
}
