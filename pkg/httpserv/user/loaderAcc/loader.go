package loaderAcc

import (
	. "GameLoaders/pkg/businesslogic/loader"
	. "GameLoaders/pkg/httpserv/user/account"
)

type User struct {
	*Account `json:"account"`
	*Loader  `json:"loader"`
}

func NewUser(account *Account, loader *Loader) *User {
	return &User{
		Account: account,
		Loader:  loader,
	}
}
