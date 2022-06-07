package customer

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/businesslogic/wallet"
)

//type Customer struct {
//	IWallet
//	sync.RWMutex
//	*account.Account
//	tasks   []ITask
//	loaders []ILoader
//}

type Model struct {
	Account *account.Model  `json:"account"`
	Wallet  *wallet.Model   `json:"wallet"`
	Tasks   []*task.Task    `json:"tasks"`
	Loaders []*loader.Model `json:"loaders"`
}

func (c *Customer) ToModel() *Model {
	tasks := make([]*task.Task, 0, len(c.tasks))
	for _, v := range c.tasks {
		if ok := v.HasMoved(); !ok {
			tasks = append(tasks, v)
		}
	}

	loaders := make([]*loader.Model, len(c.loaders))
	for i := range c.loaders {
		loaders = append(loaders, c.loaders[i].ToModel())
	}
	return &Model{
		Account: c.Account.ToModel(),
		Wallet:  c.IWallet.ToModel(),
		Loaders: loaders,
		Tasks:   tasks,
	}
}
