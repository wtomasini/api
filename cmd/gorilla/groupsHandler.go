package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
	"github.com/willTomasini/api/pkg/groups"
	"net/http"
)

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
