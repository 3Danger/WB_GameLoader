package main

import (
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/httpserv/user/account"
	"GameLoaders/pkg/httpserv/user/loaderAcc"
	"encoding/json"
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
		loader.NewLoader("Vasa Mauro"),
		loader.NewLoader("Petr Perviy"),
		loader.NewLoader("Ivan Vasiliev"),
		loader.NewLoader("Steve Jack"),
		loader.NewLoader("James Bond"),
	}
}
func GenerateTasks() []*task.Task {
	return []*task.Task{
		task.NewTask("Mac", 30),
		task.NewTask("Bananas", 40),
		task.NewTask("Bricks", 80),
		task.NewTask("Brads", 10),
	}
}

func main() {
	loader := loaderAcc.NewUser(account.NewAccount("123", "loaderee", "loader", "qwe", false), loader.NewLoader("loader"))
	fmt.Println(loader.Loader.CanMoveWeight())
	a, _ := json.Marshal(loader.Loader)
	fmt.Println(string(a))
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
	viper.AddConfigPath("configs")
	viper.SetConfigFile("server")
	return viper.ReadInConfig()
}
