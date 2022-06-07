package wallet

type Model struct {
	Money float32 `json:"money"`
}

func (w *Wallet) ToModel() *Model {
	return &Model{w.GetInfo()}
}
