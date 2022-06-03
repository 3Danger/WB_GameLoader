package main

import (
	"GameLoaders/pkg/loader"
	"GameLoaders/pkg/task"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var ok error
	alkash := loader.NewLoader()
	haha := task.NewTaskRand("haha")
	for !haha.HasMoved() && ok == nil {
		ok = alkash.Unload(haha)
	}
	if ok != nil {
		fmt.Println(ok)
	}
	if haha.HasMoved() {
		fmt.Println("Задание выполнено!")
	} else {
		fmt.Println("Задание провалено!")
	}
	fmt.Println(haha.Weight())
	fmt.Println(alkash.CanMoveWeight())
}
