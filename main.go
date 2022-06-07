package main

import (
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/httpserv/handler"
	"GameLoaders/pkg/httpserv/server"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}
func GenerateLoaders() []*loader.Loader {
	return []*loader.Loader{
		loader.NewLoaderRand("Vasa Mauro"),
		loader.NewLoaderRand("Petr Perviy"),
		loader.NewLoaderRand("Ivan Vasiliev"),
		loader.NewLoaderRand("Steve Jack"),
		loader.NewLoaderRand("James Bond"),
	}
}
func GenerateTasks() []*task.Task {
	return []*task.Task{
		{"Mac", 30},
		{"Bananas", 40},
		{"Bricks", 80},
		{"Brads", 10},
	}
}

func main() {
	if ok := initConfig(); ok != nil {
		log.Fatalln(ok)
	}
	oper := handler.NewOperator()
	//route := new(http.ServeMux)
	//route.HandleFunc("/login", oper.Login)
	serv := new(server.Server)

	ok := serv.Run(viper.GetString("port"), oper.GetRoute())
	if ok != nil {
		log.Fatalln(ok)
	}

	//loader := loaderAcc.NewUser(account.NewAccount("123", "loaderee", "loader", "qwe", false), loader.NewLoaderRand("loader"))
	//fmt.Println(loader.Loader.CanMoveWeight())
	//a, _ := json.Marshal(loader.Loader)
	//fmt.Println(string(a))
}

func main2() {
	if ok := initConfig(); ok != nil {
		log.Fatalln(ok)
	}
	tasks := GenerateTasks()
	loaders := GenerateLoaders()
	client := customer.NewCustomer(100.000, "Client one")
	client2 := customer.NewCustomer(20.000, "Client two")
	for i, v := range tasks {
		if i%2 == 0 {
			client.AddTask(v)
		} else {
			client2.AddTask(v)
		}
	}
	for i, v := range loaders {
		if i%2 == 0 {
			if ok := client.HireLoader(v); ok != nil {
				fmt.Println(ok)
				break
			}
		} else {
			if ok := client2.HireLoader(v); ok != nil {
				fmt.Println(ok)
				break
			}
		}
	}
	if ok := client.Start(); ok != nil {
		log.Println(ok)
	} else {
		log.Println("success")
	}
	if ok := client2.Start(); ok != nil {
		log.Println(ok)
	} else {
		log.Println("success")
	}
}

func initConfig() (ok error) {
	viper.AddConfigPath("./configs/")
	viper.SetConfigType("yml")
	viper.SetConfigName("server")
	return viper.ReadInConfig()
}
