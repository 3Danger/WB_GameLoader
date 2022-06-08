package customer

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/businesslogic/wallet"
	"errors"
	"math/rand"
	"sync"
)

type Customer struct {
	*wallet.Wallet
	sync.RWMutex
	*account.Account
	id      int
	tasks   []*task.Task
	loaders []*loader.Loader
}

func (c *Customer) Id() int      { return c.id }
func (c *Customer) SetId(id int) { c.id = id }

func (c *Customer) Loaders() []*loader.Loader {
	return c.loaders
}

func (c *Customer) Tasks() []*task.Task {
	return c.tasks
}

func (c *Customer) AddTask(task *task.Task) *Customer {
	c.Lock()
	c.tasks = append(c.tasks, task)
	c.Unlock()
	return c
}

func NewCustomer(account *account.Account, money float32) *Customer {
	return &Customer{
		Wallet:  wallet.NewWallet(money),
		Account: account,
		tasks:   make([]*task.Task, 0, 10),
	}
}

func NewCustomerRand(account *account.Account) *Customer {
	return &Customer{
		Wallet:  wallet.NewWallet(rand.Float32()*90_000 + 10_000),
		Account: account,
		tasks:   make([]*task.Task, 0, 10),
	}
}

func (c *Customer) Start() (ok error) {
	var okLoader error
	loaders := c.loaders
	if len(c.tasks) == 0 {
		return errors.New("there is no task")
	}
	chainTasks := new(chainOfTaskBuilder).Add(c.tasks...).Build()
	for _, ldr := range loaders {
		for okLoader == nil {
			okLoader = ldr.Unload(chainTasks.Task)
		}
		okLoader = nil
		if ok = c.SendTo(ldr, ldr.Salary()); ok != nil {
			return ok
		}
	}
	if chainTasks.HasMoved() {
		return nil
	}
	return errors.New("last task \"" + chainTasks.Name + "\" failed!")
}

func (c *Customer) HireLoader(loaders *loader.Loader) {
	c.Lock()
	c.loaders = append(c.loaders, loaders)
	c.Unlock()
}
