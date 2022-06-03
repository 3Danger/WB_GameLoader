package customer

import (
	"GameLoaders/pkg/task"
	"GameLoaders/pkg/wallet"
)

type Customer struct {
	wallet.IWallet
	tasks []*task.Task
	//loaders []*Loader
}

//func HireLoaderFromSlice()
