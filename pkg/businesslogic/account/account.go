package account

type IAccount interface {
	Id() int
	Name() string
	Login() string
	Password() string
}

type Account struct {
	id       int
	name     string
	login    string
	password string
}

func NewAccount(id int, name, username, password string) *Account {
	return &Account{id, name, username, password}
}

func (a *Account) SetId(id int) {
	a.id = id
}

func (a Account) Id() int {
	return a.id
}

func (a Account) Name() string {
	return a.name
}

func (a Account) Login() string {
	return a.login
}

func (a Account) Password() string {
	return a.password
}
