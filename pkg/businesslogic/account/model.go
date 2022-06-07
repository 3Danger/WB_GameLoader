package account

func (a *Account) ToModelAccount() *Model {
	return &Model{a.name, a.username, a.password, false}
}

type Model struct {
	Name       string `json:"name"`
	Username   string `json:"login"`
	Password   string `json:"password"`
	IsCustomer bool   `json:"is_customer"`
}
