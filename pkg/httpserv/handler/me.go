package handler

import "net/http"

func (o *Operator) Me(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
		return
	}
	acc, ok := parseToken(r.Header.Get("Authorization"))
	if ok != nil {
		writeError(w, ok.Error(), http.StatusForbidden)
		return
	}

	if !o.HasLogin(acc.Login) {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	writeData(w, o.GetUser(acc.Login).ToModel(), http.StatusOK)
}
