package handler

import (
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"net/http"
	"reflect"
)

func (o *Operator) Start(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Status method not allowed", http.StatusMethodNotAllowed)
		return
	}
	acc, ok := parseToken(r.Header.Get("Authorization"))
	if ok != nil {
		writeError(w, ok.Error(), http.StatusForbidden)
		return
	}

	iaccount := o.GetUser(acc.Login)
	if reflect.ValueOf(iaccount).IsNil() {
		writeError(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	cstmr, isCustomer := iaccount.(*customer.Customer)
	if !isCustomer {
		writeError(w, "forbidden for loaders", http.StatusForbidden)
		return
	}

	//TODO remove this кастыль
	//START TEST
	for _, l := range o.loaders {
		cstmr.HireLoader(l.(*loader.Loader))
		ok = o.db.UpdateLoader(l.(*loader.Loader), cstmr.Id())
	}
	//END TEST

	if ok = cstmr.Start(); ok != nil {
		writeResult(w, "game over")
	} else {
		writeResult(w, "you are WIN!")
	}
	ok = o.db.UpdateCustomer(cstmr)
	for _, l := range cstmr.Loaders() {
		ok = o.db.UpdateLoader(l, cstmr.Id())
		for _, t := range l.Tasks() {
			ok = o.db.UpdateTask(t, l.Id())
		}
	}
}
