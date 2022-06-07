package customer

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/interfaces"
	"GameLoaders/pkg/businesslogic/wallet"
	"errors"
	"math/rand"
	"sync"
)

type Customer struct {
	interfaces.IWallet
	sync.RWMutex
	*account.Account
	tasks   []interfaces.ITask
	loaders []interfaces.ILoader
}

func (c *Customer) Tasks() []interfaces.ITask {
	return c.tasks
}

func (c *Customer) AddTask(task interfaces.ITask) *Customer {
	c.Lock()
	c.tasks = append(c.tasks, task)
	c.Unlock()
	return c
}

func NewCustomer(account *account.Account, money float32) *Customer {
	return &Customer{
		IWallet: wallet.NewWallet(money),
		Account: account,
		tasks:   make([]interfaces.ITask, 0, 10),
	}
}

func NewCustomerRand(account *account.Account) *Customer {
	return &Customer{
		IWallet: wallet.NewWallet(rand.Float32()*90_000 + 10_000),
		Account: account,
		tasks:   make([]interfaces.ITask, 0, 10),
	}
}

func (c *Customer) Start() (ok error) {
	var okLoader error
	loaders := c.loaders
	if len(c.tasks) == 0 {
		return errors.New("there is no task")
	}
	chainTasks := new(chainOfTaskBuilder).Add(c.tasks...).Build()
	for _, loader := range loaders {
		for okLoader == nil {
			okLoader = loader.Unload(chainTasks)
		}
		okLoader = nil
		if ok = c.SendTo(loader.Salary(), loader); ok != nil {
			return ok
		}
	}
	if chainTasks.HasMoved() {
		return nil
	}
	return errors.New("last task \"" + chainTasks.GetName() + "\" failed!")
}

func (c *Customer) HireLoader(loaders interfaces.ILoader) (ok error) {
	c.Lock()
	c.loaders = append(c.loaders, loaders)
	c.Unlock()
	return nil
}
