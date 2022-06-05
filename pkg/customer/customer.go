package customer

import (
	. "GameLoaders/pkg/_interfaces"
	. "GameLoaders/pkg/wallet"
	"errors"
)

type Customer struct {
	IWallet
	name    string
	tasks   []ITask
	loaders []ILoader
}

func (c *Customer) AddTask(task ITask) *Customer {
	c.tasks = append(c.tasks, task)
	return c
}

func NewCustomer(money float32, name string) *Customer {
	return &Customer{
		IWallet: NewWallet(money),
		name:    name,
		tasks:   make([]ITask, 0, 10),
	}
}

func (c *Customer) Start() (ok error) {
	var okLoader error
	loaders := c.loaders
	chainTasks := new(ChainOfTaskBuilder).Add(c.tasks...).Build()
	for _, l := range loaders {
		for okLoader == nil {
			okLoader = l.Unload(chainTasks)
		}
		okLoader = nil
		if ok = c.SendTo(l.Salary(), l); ok != nil {
			return ok
		}
	}
	if chainTasks.HasMoved() {
		return nil
	}
	return errors.New("last task failed!")
}

func (c *Customer) HireLoader(loaders ILoader) (ok error) {
	c.loaders = append(c.loaders, loaders)
	return nil
}
