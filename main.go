package main

import (
	"GameLoaders/pkg/httpserv/database"
	"GameLoaders/pkg/httpserv/handler"
	"GameLoaders/pkg/httpserv/server"
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
	if ok := initConfig(); ok != nil {
		log.Fatalln(ok)
	}
}

func initConfig() (ok error) {
	viper.AddConfigPath("./configs/")
	viper.SetConfigType("yml")
	viper.SetConfigName("server")
	return viper.ReadInConfig()
}

func cleanDataBaseDEBUG() {
	db := database.NewDB()
	defer db.Close()
	var row *pgx.Rows
	row, _ = db.Connection().Query("DELETE FROM tasks")
	row.Close()
	row, _ = db.Connection().Query("DELETE FROM loader")
	row.Close()
	row, _ = db.Connection().Query("DELETE FROM customer")
	row.Close()
	row, _ = db.Connection().Query("DELETE FROM account")
	row.Close()
}

func main() {
	var ok error
	//cleanDataBaseDEBUG()
	db := database.NewDB()
	defer db.Close()
	op := handler.NewOperator(db)
	ok = (&server.Server{}).Run(viper.GetString("port"), op.GetRoute())
	if ok != nil {
		log.Fatalln(ok)
	}
}
