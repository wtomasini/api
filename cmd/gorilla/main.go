package gorilla

import (
	"github.com/willTomasini/api/pkg/groups"
	"github.com/willTomasini/api/pkg/users"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	userStore := users.NewMemStore()
	usersHandler := NewUsersHandler(userStore)

	groupStore := groups.NewMemStore()
	groupsHandler := NewGroupsHandler(groupStore)

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

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("500 Internal Server Error"))
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("404 Not Found"))
}
