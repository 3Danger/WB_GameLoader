package customerAcc

import (
	. "GameLoaders/pkg/businesslogic/customer"
	. "GameLoaders/pkg/httpserv/user/account"
)

type taskModel struct {
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
}

type userModel struct {
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Tasks    []taskModel `json:"tasks"`
}

type User struct {
	*Account  `json:"account"`
	*Customer `json:"customer"`
}

func NewUser(account *Account, customer *Customer) *User {
	return &User{
		Account:  account,
		Customer: customer,
	}
}

func (u *User) ToModel() interface{} {
	model := userModel{
		Name:     u.Name(),
		Username: u.Login(),
		Tasks:    nil,
	}
	for _, v := range u.Tasks() {
		model.Tasks = append(model.Tasks, taskModel{
			Name:   v.GetName(),
			Weight: v.GetWeight(),
		})
	}
	return model
}
