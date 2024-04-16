package main

import (
	"github.com/gorilla/mux"
	"github.com/willTomasini/api/controller"
	"net/http"
)

func main() {
	userStore := controller.NewUserMemStore()
	usersHandler := controller.NewUsersHandler(userStore)

	groupStore := controller.NewGroupMemStore()
	groupsHandler := controller.NewGroupsHandler(groupStore)

	home := homeHandler{}

	router := mux.NewRouter()
	router.HandleFunc("/", home.ServeHTTP)

	u := router.PathPrefix("/users").Subrouter()

	u.HandleFunc("/", usersHandler.ListUsers).Methods(http.MethodGet)
	u.HandleFunc("/", usersHandler.CreateUser).Methods(http.MethodPost)
	u.HandleFunc("/{id}", usersHandler.GetUser).Methods(http.MethodGet)
	u.HandleFunc("/{id}", usersHandler.UpdateUser).Methods(http.MethodPut)
	u.HandleFunc("/{id}", usersHandler.DeleteUser).Methods(http.MethodDelete)

	g := router.PathPrefix("/groups").Subrouter()

	g.HandleFunc("/", groupsHandler.ListGroups).Methods(http.MethodGet)
	g.HandleFunc("/", groupsHandler.CreateGroup).Methods(http.MethodPost)
	g.HandleFunc("/{id}", groupsHandler.GetGroup).Methods(http.MethodGet)
	g.HandleFunc("/{id}", groupsHandler.UpdateGroup).Methods(http.MethodPut)
	g.HandleFunc("/{id}", groupsHandler.DeleteGroup).Methods(http.MethodDelete)

	_ = http.ListenAndServe(":8010", router)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Why tho\n -Ryan"))
}
