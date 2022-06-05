package main

import (
	"GameLoaders/pkg/customer"
	"GameLoaders/pkg/loader"
	"GameLoaders/pkg/task"
	"fmt"
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
	tasks := GenerateTasks()
	loaders := GenerateLoaders()
	client := customer.NewCustomer(100.000, "Client")
	for _, v := range tasks {
		client.AddTask(v)
	}
	for _, v := range loaders {
		if ok := client.HireLoader(v); ok != nil {
			fmt.Println(ok)
			break
		}
	}
	if ok := client.Start(); ok != nil {
		log.Println(ok)
	} else {
		log.Println("success")
	}
}
