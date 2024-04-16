package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	home := homeHandler{}

	router := mux.NewRouter()
	router.HandleFunc("/", home.ServeHTTP)

	_ = http.ListenAndServe(":8010", router)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Why tho\n -Ryan"))
}
