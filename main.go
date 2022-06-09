package main

import (
	"GameLoaders/pkg/httpserv/database"
	"GameLoaders/pkg/httpserv/handler"
	"GameLoaders/pkg/httpserv/server"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var ok error
	db := database.NewDB()
	defer db.Close()

	//var row *pgx.Rows
	//row, ok = db.Connection().Query("DELETE FROM tasks")
	//row.Close()
	//row, ok = db.Connection().Query("DELETE FROM loader")
	//row.Close()
	//row, ok = db.Connection().Query("DELETE FROM customer")
	//row.Close()
	//row, ok = db.Connection().Query("DELETE FROM account")
	//row.Close()

	op := handler.NewOperator(db)
	ok = (&server.Server{}).Run("8080", op.GetRoute())
	if ok != nil {
		log.Fatalln(ok)
	}
}

func initConfig() (ok error) {
	viper.AddConfigPath("./configs/")
	viper.SetConfigType("yml")
	viper.SetConfigName("server")
	return viper.ReadInConfig()
}
