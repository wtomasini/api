package users

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type MemStore struct {
	list map[string]User
}

func NewMemStore() *MemStore {
	list := make(map[string]User)
	return &MemStore{
		list,
	}
}

func (m MemStore) Add(name string, User User) error {
	m.list[name] = User
	return nil
}

func (m MemStore) Get(name string) (User, error) {

	if val, ok := m.list[name]; ok {
		return val, nil
	}

	return User{}, ErrNotFound
}

func (m MemStore) List() (map[string]User, error) {
	return m.list, nil
}

func (m MemStore) Update(name string, User User) error {

	if _, ok := m.list[name]; ok {
		m.list[name] = User
		return nil
	}

	return ErrNotFound
}

func (m MemStore) Remove(name string) error {
	delete(m.list, name)
	return nil
}
