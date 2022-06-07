package loader

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/task"
	"GameLoaders/pkg/businesslogic/wallet"
)

type Model struct {
	Account        *account.Model        `json:"account"`
	Wallet         *wallet.Model         `json:"wallet"`
	SuccessTasks   map[string]*task.Task `json:"success_tasks"`
	MaxWeightTrans float32               `json:"max_weight_trans"`
	Salary         float32               `json:"salary"`
	Fatigue        float32               `json:"fatigue"`
	Drunk          bool                  `json:"drunk"`
}

func (l *Loader) ToModel() interface{} {
	tasks := make(map[string]*task.Task, 0)
	for _, v := range l.tasks {
		if v.HasMoved() {
			if _, ok := tasks[v.GetName()]; !ok {
				tasks[v.GetName()] = v
			}
		}
	}
	return &Model{
		Account:        l.Account.ToModelAccount(),
		Wallet:         l.Wallet.ToModel(),
		SuccessTasks:   tasks,
		MaxWeightTrans: l.maxWeightTrans,
		Salary:         l.salary,
		Fatigue:        l.fatigue,
		Drunk:          l.drunk,
	}
}
