package database

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"fmt"
)

func (d *DB) BindTaskTo(task *task.Task, acc *account.Account) (ok error) {
	q := fmt.Sprintf("UPDATE tasks SET account_id = %d WHERE id = %s", acc.Id(), task.Id)
	_, ok = d.connect.Query(q)
	return ok
}

func (d *DB) BindLoaderToCustomer(l *loader.Loader, c *customer.Customer) (ok error) {
	q := fmt.Sprintf("UPDATE loader SET customer_id = %d WHERE id = %d",
		c.Id(), l.Id())
	_, ok = d.connect.Query(q)
	return ok
}
