package controller

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

// I'm copying the logic over for users and groups, but there has to be a better way of doing this,
// by handling both with the same function.

type MemStore struct {
	list map[string]Group
}

func NewMemStore() *MemStore {
	list := make(map[string]Group)
	return &MemStore{
		list,
	}
}

func (m MemStore) Add(name string, group Group) error {
	m.list[name] = group
	return nil
}

func (m MemStore) Get(name string) (Group, error) {

	if val, ok := m.list[name]; ok {
		return val, nil
	}

	return Group{}, ErrNotFound
}

func (m MemStore) List() (map[string]Group, error) {
	return m.list, nil
}

func (m MemStore) Update(name string, group Group) error {

	if _, ok := m.list[name]; ok {
		m.list[name] = group
		return nil
	}

	return ErrNotFound
}

func (m MemStore) Remove(name string) error {
	delete(m.list, name)
	return nil
}
