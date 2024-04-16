package controller

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

// I'm copying the logic over for users and groups, but there has to be a better way of doing this,
// by handling both with the same functions.

type UserMemStore struct {
	list map[string]User
}

func NewUserMemStore() *UserMemStore {
	list := make(map[string]User)
	return &UserMemStore{
		list,
	}
}

func (m UserMemStore) Add(name string, user User) error {
	m.list[name] = user
	return nil
}

func (m UserMemStore) Get(name string) (User, error) {

	if val, ok := m.list[name]; ok {
		return val, nil
	}

	return User{}, ErrNotFound
}

func (m UserMemStore) List() (map[string]User, error) {
	return m.list, nil
}

func (m UserMemStore) Update(name string, user User) error {

	if _, ok := m.list[name]; ok {
		m.list[name] = user
		return nil
	}

	return ErrNotFound
}

func (m UserMemStore) Remove(name string) error {
	delete(m.list, name)
	return nil
}

type GroupMemStore struct {
	list map[string]Group
}

func NewGroupMemStore() *GroupMemStore {
	list := make(map[string]Group)
	return &GroupMemStore{
		list,
	}
}

func (m GroupMemStore) Add(name string, group Group) error {
	m.list[name] = group
	return nil
}

func (m GroupMemStore) Get(name string) (Group, error) {

	if val, ok := m.list[name]; ok {
		return val, nil
	}

	return Group{}, ErrNotFound
}

func (m GroupMemStore) List() (map[string]Group, error) {
	return m.list, nil
}

func (m GroupMemStore) Update(name string, group Group) error {

	if _, ok := m.list[name]; ok {
		m.list[name] = group
		return nil
	}

	return ErrNotFound
}

func (m GroupMemStore) Remove(name string) error {
	delete(m.list, name)
	return nil
}
