package user

import "sync"

type Repository interface {
	FindAll() ([]User, error)
	Save(user User) error
	FindById(id string) (User, error)
}

type memoryRepository struct {
	mu    sync.RWMutex
	users []User
}

func NewRepository() Repository {
	return &memoryRepository{
		users: []User{},
	}
}

func (r *memoryRepository) FindById(id string) (User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, ErrUserNotFound
}

func (r *memoryRepository) FindAll() ([]User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.users, nil
}

func (r *memoryRepository) Save(user User) error {
	if user.ID == "" || user.Name == "" {
		return ErrInvalidUser
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.users = append(r.users, user)
	return nil
}
