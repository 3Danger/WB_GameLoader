package loader

type Model struct {
	Id    int    `json:"id"`
	Login string `json:"login"`

	MaxWeight float32 `json:"max_weight"`
	Money     float32 `json:"money"`
	Drunk     bool    `json:"drunk"`
	Fatigue   float32 `json:"fatigue"`
}

//ToModel показать свои характеристики (вес, деньги, пьянство, усталость)
func (l *Loader) ToModel() interface{} {
	return &Model{
		Id:        l.Id(),
		Login:     l.Login(),
		MaxWeight: l.MaxWeightTrans(),
		Money:     l.Wallet.GetInfo(),
		Drunk:     l.Drunk(),
		Fatigue:   l.Fatigue(),
	}
}
