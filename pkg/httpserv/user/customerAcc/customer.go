package customerAcc

import (
	. "GameLoaders/pkg/businesslogic/customer"
	. "GameLoaders/pkg/httpserv/user/account"
)

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
