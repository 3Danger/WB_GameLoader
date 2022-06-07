package handler

import (
	"GameLoaders/pkg/businesslogic/account"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io"
	"strings"
)

type ClaimsAccount struct {
	jwt.StandardClaims `json:"-"`
	account.Model      `json:"account.model"`
}

func accountParseFrom(body io.Reader) (accClaims *ClaimsAccount, ok error) {
	accClaims = new(ClaimsAccount)
	accModel := account.Model{}
	ok = json.NewDecoder(body).Decode(&accModel)
	if ok == nil {
		if accModel.Username == "" {
			ok = errors.New("invalid username")
		}
		if accModel.Password == "" {
			ok = errors.New("invalid password")
		}
	}
	if ok != nil {
		return nil, ok
	}
	accClaims.Model = accModel
	return accClaims, nil
}

func (a *ClaimsAccount) generateToken() (token string, ok error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, a).SignedString([]byte(signingKey))
}

func parseToken(accessTokenBeaver string) (acc *ClaimsAccount, ok error) {
	var accessToken string
	var isCut bool
	acc = new(ClaimsAccount)

	if _, accessToken, isCut = strings.Cut(accessTokenBeaver, "Bearer "); !isCut {
		return nil, errors.New("token invalid")
	}
	token, ok := jwt.ParseWithClaims(accessToken, &ClaimsAccount{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if ok != nil {
		return nil, ok
	}

	acc, okay := token.Claims.(*ClaimsAccount)
	if !okay {
		return nil, errors.New("token claims are not of type *account")
	}
	return acc, nil
}
