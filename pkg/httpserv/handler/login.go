package handler

import (
	_ "github.com/dgrijalva/jwt-go"
	"net/http"
)

type Token struct {
	DataOfToken string `json:"token"`
}

func (o *Operator) Login(w http.ResponseWriter, r *http.Request) {
	var ok error
	var acc *ClaimsAccount

	if r.Method != http.MethodPost {
		writeError(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	if acc, ok = accountParseFrom(r.Body); ok != nil {
		writeError(w, ok.Error(), http.StatusBadRequest)
		return
	}

	if user := o.GetUser(acc.Login); user == IAccount(nil) {
		writeError(w, "account: "+acc.Login+" not found", http.StatusBadRequest)
	} else if generatePasswordHash(acc.Password) != user.Password() {
		writeError(w, "password invalid", http.StatusBadGateway)
	} else {
		sign, _ := acc.generateToken()
		writeData(w, Token{sign}, http.StatusOK)
	}
}
