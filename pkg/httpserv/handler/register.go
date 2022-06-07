package handler

import (
	"GameLoaders/pkg/businesslogic/account"
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func (o *Operator) Register(w http.ResponseWriter, r *http.Request) {
	var ok error
	var acc *ClaimsAccount

	if http.MethodPost != r.Method {
		writeError(w, "bad method", http.StatusMethodNotAllowed)
		return
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
		o.Add(customer.NewCustomerRand(account.NewAccountFromModel(&acc.Model)))
	} else {
		o.Add(loader.NewLoaderRand(account.NewAccountFromModel(&acc.Model)))
	}
	writeResult(w, "success")
}
