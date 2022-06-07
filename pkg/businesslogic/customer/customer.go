package customer

import (
	. "GameLoaders/pkg/businesslogic/_interfaces"
	. "GameLoaders/pkg/businesslogic/wallet"
	"errors"
	"math/rand"
	"sync"
)

type Customer struct {
	IWallet
	sync.RWMutex
	name    string
	tasks   []ITask
	loaders []ILoader
}

func (c *Customer) Tasks() []ITask {
	return c.tasks
}

func (c *Customer) AddTask(task ITask) *Customer {
	c.Lock()
	c.tasks = append(c.tasks, task)
	c.Unlock()
	return c
}

func NewCustomer(money float32, name string) *Customer {
	return &Customer{
		IWallet: NewWallet(money),
		name:    name,
		tasks:   make([]ITask, 0, 10),
	}
}

func NewCustomerRand(name string) *Customer {
	return &Customer{
		IWallet: NewWallet(rand.Float32()*90_000 + 10_000),
		name:    name,
		tasks:   make([]ITask, 0, 10),
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

func (c *Customer) HireLoader(loaders ILoader) (ok error) {
	c.Lock()
	c.loaders = append(c.loaders, loaders)
	c.Unlock()
	return nil
}
