package handler

import (
	"GameLoaders/pkg/businesslogic/task"
	"encoding/json"
	"net/http"
)

type TasksModel struct {
	Tasks []*task.Task `json:"tasks"`
}

func (o *Operator) Tasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		t := TasksModel{make([]*task.Task, 0)}
		if ok := json.NewDecoder(r.Body).Decode(&t); ok != nil {
			writeError(w, ok.Error(), http.StatusBadRequest)
		} else {
			if len(o.customers) == 0 {
				writeError(w, "customers not found", http.StatusConflict)
				return
			}
			o.AddTasks(t.Tasks...)
			writeResult(w, "success")
		}
		return
	}
	if r.Method == http.MethodGet {
		acc, ok := parseToken(r.Header.Get("Authorization"))
		if ok != nil {
			writeError(w, ok.Error(), http.StatusBadRequest)
			return
		}
		writeData(w, o.GetUser(acc.Login).Tasks(), http.StatusOK)
		return
	}
}
