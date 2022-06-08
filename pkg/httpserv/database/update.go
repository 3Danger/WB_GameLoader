package database

import (
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"github.com/jackc/pgx"
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
	var rows *pgx.Rows
	q := "UPDATE customer SET money = $1 WHERE id = $2"
	rows, ok = d.connect.Query(q, c.Wallet.GetInfo(), c.Id())
	rows.Close()
	return ok
}

func (d *DB) UpdateLoader(l *loader.Loader, customer_id int) (ok error) {
	var rows *pgx.Rows
	q := "UPDATE loader SET customer_id = $1, money = $2, fatigue = $3 WHERE id = $4"
	rows, ok = d.connect.Query(q, customer_id, l.Wallet.GetInfo(), l.Fatigue(), l.Id())
	rows.Close()
	return ok
}

func (d *DB) UpdateTask(t *task.Task, account_id int) (ok error) {
	var rows *pgx.Rows
	q := "UPDATE loader SET weight = $1, account_id = $2 WHERE id = $3"
	rows, ok = d.connect.Query(q, t.Weight, account_id, t.Id)
	rows.Close()
	return ok
}
