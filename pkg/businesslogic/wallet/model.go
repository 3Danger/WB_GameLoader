package wallet

type Model struct {
	Money float32 `json:"money"`
}

func (w *Wallet) ToModel() interface{} {
	return Model{w.GetInfo()}
}
