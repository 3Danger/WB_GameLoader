package account

type IAccount interface {
	GetName() string
	GetUserName() string
	GetPassword() string
	ToModel() Model
}

type Account struct {
	name     string
	username string
	password string
}

func NewAccount(name, username, password string) *Account {
	return &Account{name, username, password}
}

func NewAccountFromModel(model *Model) *Account {
	return &Account{model.Name, model.Username, model.Password}
}

func (a Account) GetName() string {
	return a.name
}

func (a Account) GetUserName() string {
	return a.username
}

func (a Account) GetPassword() string {
	return a.password
}
