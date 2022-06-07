package account

func (a *Account) ToModel() Model {
	return Model{a.name, a.username, a.password}
}

type Model struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
