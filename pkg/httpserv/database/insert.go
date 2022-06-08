package database

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"github.com/jackc/pgx"
	"log"
)

func (d *DB) InsertLoader(l *loader.Loader) (ok error) {
	var tx *pgx.Tx
	var id int
	if tx, ok = d.connect.Begin(); ok != nil {
		return ok
	}
	if id, ok = d.insertAccount(tx, l.Account); ok != nil {
		return rollback(tx, ok)
	}
	l.Account.SetId(id)
	if id, ok = d.insertLoader(tx, l); ok != nil {
		return rollback(tx, ok)
	}
	return nil
}

func (d *DB) InsertCustomer(c *customer.Customer) (ok error) {
	var (
		tx *pgx.Tx
		id int
	)
	if tx, ok = d.connect.Begin(); ok != nil {
		return ok
	}
	if id, ok = d.insertAccount(tx, c.Account); ok != nil {
		return rollback(tx, ok)
	}
	c.Account.SetId(id)
	if id, ok = d.insertCustomer(tx, c); ok != nil {
		return rollback(tx, ok)
	}
	c.SetId(id)
	if ok = tx.Commit(); ok != nil {
		return rollback(tx, ok)
	}
	return ok
}

func (d *DB) InsertTask(t *task.Task, account_id int) (ok error) {
	var query string
	var rows *pgx.Rows

	query = `INSERT INTO tasks (account_id, name, weight) VALUES ($1, $2, $3)`
	rows, ok = d.connect.Query(query, account_id, t.Name, t.Weight)
	rows.Close()
	return ok
}

func rollback(tx *pgx.Tx, ok error) error {
	if okRoll := tx.Rollback(); okRoll != nil {
		log.Fatalln(okRoll)
	}
	return ok
}

func (d *DB) insertAccount(tx *pgx.Tx, a *account.Account) (id int, ok error) {
	var (
		rows     *pgx.Rows
		queryStr string
	)
	queryStr = "INSERT INTO account (login, password) VALUES ($1, $2) RETURNING id"
	if rows, ok = tx.Query(queryStr, a.Login(), a.Password()); ok != nil {
		return 0, ok
	}
	defer rows.Close()
	rows.Next()
	if ok = rows.Scan(&id); ok != nil {
		return 0, ok
	}
	return id, nil
}

func (d *DB) insertCustomer(tx *pgx.Tx, c *customer.Customer) (id int, ok error) {
	var (
		rows     *pgx.Rows
		queryStr string
	)
	queryStr = "INSERT INTO customer (account_id, money) VALUES ($1, $2) RETURNING id"
	if rows, ok = tx.Query(queryStr, c.Account.Id(), c.Wallet.GetInfo()); ok != nil {
		return 0, ok
	}
	defer rows.Close()
	rows.Next()
	if ok = rows.Scan(&id); ok != nil {
		return 0, ok
	}
	return id, nil
}

func (d *DB) insertLoader(tx *pgx.Tx, l *loader.Loader) (id int, ok error) {
	var rows *pgx.Rows
	var queryStr string

	queryStr = "INSERT INTO loader (" +
		"account_id, customer_id, money, salary, max_weight_trans, fatigue, drunk" +
		") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING loader.id"
	if rows, ok = tx.Query(queryStr,
		l.Account.Id(),
		0,
		l.Wallet.GetInfo(),
		l.Salary(),
		l.MaxWeightTrans(),
		l.Fatigue(),
		l.Drunk(),
	); ok != nil {
		return 0, ok
	}
	defer rows.Close()
	rows.Next()
	if ok = rows.Scan(&id); ok != nil {
		return 0, ok
	}
	return id, nil
}
