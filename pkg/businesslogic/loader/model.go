package loader

type Model struct {
	Account        interface{}   `json:"account"`
	Wallet         interface{}   `json:"wallet"`
	SuccessTasks   []interface{} `json:"success_tasks"`
	MaxWeightTrans float32       `json:"max_weight_trans"`
	Salary         float32       `json:"salary"`
	Fatigue        float32       `json:"fatigue"`
	Drunk          bool          `json:"drunk"`
}

func (l *Loader) ToModel() interface{} {
	tasks := make([]interface{}, 0)
	for _, v := range l.tasks {
		if v.HasMoved() {
			tasks = append(tasks, v)
		}
	}
	return Model{
		Account:        l.Account.ToModel(),
		Wallet:         l.IWallet.ToModel(),
		SuccessTasks:   tasks,
		MaxWeightTrans: l.maxWeightTrans,
		Salary:         l.salary,
		Fatigue:        l.fatigue,
		Drunk:          l.drunk,
	}
}
