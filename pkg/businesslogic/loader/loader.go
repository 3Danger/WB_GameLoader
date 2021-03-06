package loader

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/businesslogic/wallet"
	"errors"
	"math/rand"
)

type Loader struct {
	*wallet.Wallet
	*account.Account
	id                 int
	customer_id        int
	custommerAccountId int
	tasks              map[string]*task.Task
	maxWeightTrans     float32 //5-30kg
	salary             float32 //ЗП
	fatigue            float32 //усталость
	drunk              bool
}

func (l *Loader) CustomerId() int       { return l.customer_id }
func (l *Loader) SetCustomerId(cid int) { l.customer_id = cid }

func (l *Loader) CustomerAccountId() int        { return l.custommerAccountId }
func (l *Loader) SetCustomerAccountId(caid int) { l.custommerAccountId = caid }

func (l *Loader) Id() int      { return l.id }
func (l *Loader) SetId(id int) { l.id = id }

func (l *Loader) Tasks() []*task.Task {
	tasks := make([]*task.Task, 0, len(l.tasks))
	for _, v := range l.tasks {
		if v.HasMoved() {
			tasks = append(tasks, v)
		}
	}
	return tasks
}

func (l *Loader) MaxWeightTrans() float32 {
	return l.maxWeightTrans
}

func (l *Loader) Fatigue() float32 {
	return l.fatigue
}

func (l *Loader) Drunk() bool {
	return l.drunk
}

func (l *Loader) Salary() float32 {
	return l.salary
}

func NewLoaderRand(account *account.Account) *Loader {
	drunk := rand.Int()&1 == 0
	return &Loader{
		Account:        account,
		Wallet:         wallet.NewWallet(0),
		tasks:          make(map[string]*task.Task),
		maxWeightTrans: rand.Float32()*25 + 5,
		salary:         rand.Float32()*20_000 + 10_000,
		fatigue:        0.0,
		drunk:          drunk,
	}
}

func (l *Loader) CanMoveWeight() float32 {
	fatigue := l.fatigue
	if l.drunk {
		fatigue += 0.5
	}
	if fatigue > 1.0 {
		fatigue = 1.0
	}
	return l.maxWeightTrans * (1.0 - fatigue)
}

func (l *Loader) Unload(task *task.Task) error {
	//- формула рассчета переносимого веса
	// (вес*(100 - усталость/100)*(пьянство/100))
	canMoveWeight := l.CanMoveWeight()
	if canMoveWeight <= 0. {
		return errors.New(l.Login() + " is very tired")
	}
	task.Unload(canMoveWeight)
	if _, ok := l.tasks[task.Id]; !ok {
		l.tasks[task.Id] = task
	}
	l.fatigue += 0.2
	return nil
}

func (l *Loader) AddTask(tasks ...*task.Task) {
	for _, v := range tasks {
		l.tasks[v.Id] = v
	}
}
