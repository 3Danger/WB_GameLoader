package database

import (
	pgx "github.com/jackc/pgx"
	"log"
)

type DB struct {
	connect *pgx.Conn
}

func (d *DB) Connection() *pgx.Conn { return d.connect }

func NewDB() (db *DB) {
	db = new(DB)
	connect, err := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "csamuro",
		User:     "csamuro",
		Password: "qwe",
	})
	if err != nil {
		log.Fatalln(err)
	}
	db.connect = connect
	return
}

func (d *DB) Close() {
	_ = d.connect.Close()
}
