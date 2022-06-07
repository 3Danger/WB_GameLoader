package customer

import (
	"GameLoaders/pkg/businesslogic/interfaces"
)

//type Customer struct {
//	IWallet
//	sync.RWMutex
//	*account.Account
//	tasks   []ITask
//	loaders []ILoader
//}

type Model struct {
	Account interface{}          `json:"account"`
	Wallet  interface{}          `json:"wallet"`
	Tasks   []interfaces.ITask   `json:"tasks"`
	Loaders []interfaces.ILoader `json:"loaders"`
}

func (c *Customer) ToModel() interface{} {
	tasks := make([]interfaces.ITask, 0, len(c.tasks))
	for _, v := range c.tasks {
		if ok := v.HasMoved(); !ok {
			tasks = append(tasks, v)
		}
	}
	return Model{
		Account: c.Account.ToModel(),
		Wallet:  c.IWallet.ToModel(),
		Loaders: c.loaders,
		Tasks:   tasks,
	}
}
