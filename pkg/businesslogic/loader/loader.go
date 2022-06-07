package loader

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/businesslogic/wallet"
	"errors"
	"math/rand"
	"sync"
)

type Loader struct {
	*wallet.Wallet
	sync.RWMutex
	*account.Account
	tasks          map[string]*task.Task
	maxWeightTrans float32 //5-30kg
	salary         float32 //ЗП
	fatigue        float32 //усталость
	drunk          bool
}

func (l *Loader) Salary() float32 {
	l.RLock()
	defer l.RUnlock()
	return l.salary
}

func NewLoaderFromModel(model *Model) *Loader {
	return &Loader{
		Wallet:         wallet.NewWalletFromModel(model.Wallet),
		RWMutex:        sync.RWMutex{},
		Account:        account.NewAccountFromModel(model.Account),
		tasks:          model.SuccessTasks,
		maxWeightTrans: 0,
		salary:         0,
		fatigue:        0,
		drunk:          false,
	}
}

func NewLoaderRand(account *account.Account) *Loader {
	drunk := rand.Int()&1 == 0
	return &Loader{
		Account:        account,
		Wallet:         wallet.NewWallet(0),
		tasks:          make(map[string]*task.Task),
		maxWeightTrans: rand.Float32()*25 + 5,
		salary:         rand.Float32()*20 + 10,
		fatigue:        0.0,
		drunk:          drunk,
	}
}

func (l *Loader) CanMoveWeight() float32 {
	l.RLock()
	fatigue := l.fatigue
	if l.drunk {
		fatigue += 0.5
	}
	l.RUnlock()
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
		return errors.New(l.GetName() + " is very tired")
	}
	task.Unload(canMoveWeight)
	if _, ok := l.tasks[task.GetName()]; !ok {
		l.tasks[task.GetName()] = task
	}
	l.Lock()
	l.fatigue += 0.2
	l.Unlock()
	return nil
}
