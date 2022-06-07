package handler

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Error struct {
	Msg any `json:"error"`
}

type Result struct {
	Msg string `json:"result"`
}

func writeError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	defer w.WriteHeader(status)
	bt, _ := json.Marshal(Error{msg})
	if _, ok := w.Write(bt); ok != nil {
		status = 501
		log.Fatalln(ok)
	}
}

func writeResult(w http.ResponseWriter, msg string) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	var status = http.StatusOK
	defer w.WriteHeader(status)
	bt, _ := json.Marshal(Result{msg})
	if _, ok := w.Write(bt); ok != nil {
		status = 501
		log.Fatalln(ok)
	}
}

func writeData(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	if bts, ok := json.Marshal(data); ok != nil {
		status = http.StatusInternalServerError
	} else {
		if _, ok = w.Write(bts); ok != nil {
			status = http.StatusInternalServerError
		}
	}
	w.WriteHeader(status)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
