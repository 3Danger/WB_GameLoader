package account

type IAccount interface {
	Id() int
	Name() string
	Login() string
	Password() string
}

type Account struct {
	id       int
	login    string
	password string
}

func NewAccount(id int, username, password string) *Account {
	return &Account{id, username, password}
}

func (a *Account) SetId(id int) {
	a.id = id
}

func (a Account) Id() int {
	return a.id
}

func (a *Account) SetLogin(login string) { a.login = login }

func (a Account) Login() string {
	return a.login
}

func (a *Account) SetPassword(password string) { a.password = password }

func (a Account) Password() string {
	return a.password
}
