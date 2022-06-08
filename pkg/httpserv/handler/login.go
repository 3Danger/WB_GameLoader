package handler

import (
	_ "github.com/dgrijalva/jwt-go"
	"net/http"
)

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

	if !o.HasLogin(acc.Login) {
		writeError(w, "account: "+acc.Login+" not found", http.StatusBadRequest)
		return
	}
	sign, _ := acc.generateToken()

	writeData(w, struct {
		Token string `json:"token"`
	}{sign}, http.StatusOK)
}
