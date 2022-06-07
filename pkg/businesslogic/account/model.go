package account

func (a *Account) ToModel() Model {
	return Model{a.name, a.username, a.password, false}
}

type Model struct {
	Name       string `json:"name"`
	Username   string `json:"login"`
	Password   string `json:"password"`
	IsCustomer bool   `json:"is_customer"`
}
