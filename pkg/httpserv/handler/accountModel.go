package handler

import (
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	account2 "GameLoaders/pkg/httpserv/user/account"
	"GameLoaders/pkg/httpserv/user/customerAcc"
	"GameLoaders/pkg/httpserv/user/loaderAcc"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io"
	"strings"
)

type account struct {
	jwt.StandardClaims `json:"-"`
	Id                 string `json:"-"`
	Username           string `json:"login"`
	Password           string `json:"password"`
	Name               string `json:"name"`
	IsCustomer         bool   `json:"is_customer"`
}

func accountParseFrom(body io.Reader) (acc *account, ok error) {
	acc = new(account)
	ok = json.NewDecoder(body).Decode(acc)
	if ok == nil {
		if acc.Username == "" {
			ok = errors.New("invalid username")
		}
		if acc.Password == "" {
			ok = errors.New("invalid password")
		}
	}
	if ok != nil {
		return nil, ok
	}
	return acc, nil
}

func (a *account) toCustomer() *customerAcc.User {
	return customerAcc.NewUser(
		account2.NewAccount(a.Id, a.Username, a.Password, true),
		customer.NewCustomerRand(a.Name),
	)
}

func (a *account) toLoader() *loaderAcc.User {
	return loaderAcc.NewUser(
		account2.NewAccount(a.Id, a.Username, a.Password, false),
		loader.NewLoaderRand(a.Name),
	)
}

func (a *account) generateToken() (token string, ok error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, a).SignedString([]byte(signingKey))
}

func parseToken(accessTokenBeaver string) (acc *account, ok error) {
	var accessToken string
	var isCut bool
	acc = new(account)

	if _, accessToken, isCut = strings.Cut(accessTokenBeaver, "Bearer "); !isCut {
		return nil, errors.New("token invalid")
	}
	token, ok := jwt.ParseWithClaims(accessToken, &account{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if ok != nil {
		return nil, ok
	}

	acc, okay := token.Claims.(*account)
	if !okay {
		return nil, errors.New("token claims are not of type *account")
	}
	return acc, nil
}
