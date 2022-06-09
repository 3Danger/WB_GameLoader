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
	if ok := initConfig(); ok != nil {
		log.Fatalln(ok)
	}
}

func initConfig() (ok error) {
	viper.AddConfigPath("./configs/")
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	var ok error

	config := new(database.ConfigDB)
	if ok = viper.Unmarshal(config); ok != nil {
		log.Fatalln(ok)
	}
	db := database.NewDB(config)
	defer db.Close()

	op := handler.NewOperator(db)
	serv := server.NewServer(viper.GetString("serverPort"), op.GetRoute())
	if ok = serv.Run(); ok != nil {
		log.Fatalln(ok)
	}
}
