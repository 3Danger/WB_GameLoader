package handler

import (
	"GameLoaders/pkg/businesslogic/customer"
	"GameLoaders/pkg/businesslogic/loader"
	"GameLoaders/pkg/businesslogic/task"
	"net/http"
	"reflect"
)

func howMuchNeedLoaders(tasksPtrs []*task.Task, loadersPtrs []*loader.Loader) int {
	tasks := make([]task.Task, 0, len(tasksPtrs))
	loaders := make([]loader.Loader, 0, len(loadersPtrs))

	for _, t := range tasksPtrs {
		tasks = append(tasks, *t)
	}
	for _, l := range loadersPtrs {
		loaders = append(loaders, *l)
	}
	i := struct{ l, t int }{0, 0}
	for ; i.l < len(loaders); i.l++ {
		for loaders[i.l].CanMoveWeight() > 0. {
			ok := loaders[i.l].Unload(&tasks[i.t])
			if tasks[i.t].HasMoved() {
				i.t++
				if i.t == len(tasks) {
					return i.l + 1
				}
				continue
			}
			if ok != nil {
				break
			}
		}
	}
	return i.l
}

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

	loaders := make([]*loader.Loader, 0, len(o.loaders))
	for _, v := range o.loaders {
		loaders = append(loaders, v.(*loader.Loader))
	}
	howMany := howMuchNeedLoaders(cstmr.Tasks(), loaders)
	for i := 0; i < howMany; i++ {
		cstmr.HireLoader(loaders[i])
		ok = o.db.UpdateLoader(loaders[i], cstmr.Id())
	}

	if ok = cstmr.Start(); ok != nil {
		writeResult(w, "game over, "+ok.Error())
	} else {
		writeResult(w, "you are WIN!")
	}
	ok = o.db.UpdateCustomer(cstmr)
	for _, l := range cstmr.Loaders() {
		ok = o.db.UpdateLoader(l, cstmr.Id())
		for _, t := range l.Tasks() {
			ok = o.db.UpdateTask(t, l.Account.Id())
		}
	}
}
