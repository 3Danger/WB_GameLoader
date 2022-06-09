package database

import (
	pgx "github.com/jackc/pgx"
	"log"
)

type ConfigDB struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	DBName   string `yaml:"DBName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}

type DB struct {
	connect *pgx.Conn
}

func (d *DB) Connection() *pgx.Conn { return d.connect }

func NewDB(config *ConfigDB) (db *DB) {
	db = new(DB)
	connect, err := pgx.Connect(pgx.ConnConfig{
		Host:     config.Host,
		Port:     config.Port,
		Database: config.DBName,
		User:     config.UserName,
		Password: config.Password,
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
