package database

import (
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"fmt"
)

/*
customer
	money
loader
	customer_id
	money
	fatigue
task
	weight
	account_id
*/

func (d *DB) UpdateCustomer(c *customer.Customer) (ok error) {
	q := fmt.Sprintf("UPDATE customer SET money = %f WHERE id = %d",
		c.Wallet.GetInfo(), c.Id())
	_, ok = d.connect.Query(q)
	return ok
}

func (d *DB) UpdateLoader(l *loader.Loader, customer_id int) (ok error) {
	q := fmt.Sprintf("UPDATE loader SET customer_id = %d, money = %f, fatigue = %f WHERE id = %d",
		customer_id, l.Wallet.GetInfo(), l.Fatigue(), l.Id())
	_, ok = d.connect.Query(q)
	return ok
}

func (d *DB) UpdateTask(t *task.Task, account_id int) (ok error) {
	q := fmt.Sprintf("UPDATE loader SET weight = %f, account_id = %d WHERE id = %d",
		t.Weight, account_id, t.Id)
	_, ok = d.connect.Query(q)
	return ok
}
