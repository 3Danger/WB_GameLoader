package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func (o *Operator) Register(w http.ResponseWriter, r *http.Request) {
	var (
		ok  error
		acc *account
	)
	if http.MethodPost != r.Method {
		writeError(w, "bad method", http.StatusMethodNotAllowed)
		return
	}
	for _, v := range r.Cookies() {
		fmt.Println("---> Cookies:", v.Raw)
	}
	if acc, ok = accountParseFrom(r.Body); ok != nil {
		writeError(w, ok.Error(), http.StatusBadRequest)
		return
	}
	if o.HasLogin(acc.Username) {
		//TODO уточнить код ошибки
		writeError(w, "login: \""+acc.Username+"\" already use", 418)
		return
	}
	acc.Password = generatePasswordHash(acc.Password)
	acc.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	}
	if acc.IsCustomer {
		o.Add()
	}
	o.Add(acc)
	/*
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, acc)
		t, _ := token.SignedString([]byte(signingKey))
		fmt.Println("signed:", t)
		w.Header().Set("Beaver", t)
		_ = token
	*/
	//data, ok := token.SignedString([]byte(signingKey))
	//jwt.Parse
	//if ok != nil {
	//	writeError(w, ok.Error(), 501)
	//	return
	//}

	//w.Header().Add("Set-Cookie", data)
	writeResult(w, "success")
}
