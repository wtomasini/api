package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
	"github.com/willTomasini/api/pkg/groups"
	"github.com/willTomasini/api/pkg/users"
	"net/http"
)

type userStore interface {
	Add(name string, user users.User) error
	Get(name string) (users.User, error)
	List() (map[string]users.User, error)
	Update(name string, user users.User) error
	Remove(name string) error
}

type UsersHandler struct {
	store userStore
}

func NewUsersHandler(s userStore) *UsersHandler {
	return &UsersHandler{
		store: s,
	}
}

func (h UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user users.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	resourceID := slug.Make(user.Name)

	if err := h.store.Add(resourceID, user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h UsersHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	userList, err := h.store.List()

	jsonBytes, err := json.Marshal(userList)

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := h.store.Get(id)
	if err != nil {
		if errors.Is(err, users.ErrNotFound) {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.store.Update(id, user); err != nil {
		if errors.Is(err, users.ErrNotFound) {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("done with update")
	fmt.Println(string(jsonBytes))
	w.Write(jsonBytes)
}

func (h UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.store.Remove(id); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type groupStore interface {
	Add(name string, group groups.Group) error
	Get(name string) (groups.Group, error)
	List() (map[string]groups.Group, error)
	Update(name string, group groups.Group) error
	Remove(name string) error
}

type GroupsHandler struct {
	store groupStore
}

func NewGroupsHandler(s groupStore) *GroupsHandler {
	return &GroupsHandler{
		store: s,
	}
}

func (h GroupsHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group groups.Group

	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	resourceID := slug.Make(group.Name)

	if err := h.store.Add(resourceID, group); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h GroupsHandler) ListGroups(w http.ResponseWriter, r *http.Request) {
	groupList, err := h.store.List()

	jsonBytes, err := json.Marshal(groupList)

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h GroupsHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	group, err := h.store.Get(id)
	if err != nil {
		if errors.Is(err, groups.ErrNotFound) {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}
	jsonBytes, err := json.Marshal(group)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h GroupsHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var group groups.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.store.Update(id, group); err != nil {
		if errors.Is(err, groups.ErrNotFound) {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(group)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("done with update")
	fmt.Println(string(jsonBytes))
	w.Write(jsonBytes)
}

func (h GroupsHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.store.Remove(id); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("500 Internal Server Error"))
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("404 Not Found"))
}
