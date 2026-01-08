package user

import "github.com/google/uuid"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *Service) CreateUser(name string) (User, error) {
	if name == "" {
		return User{}, ErrInvalidUser
	}

	user := User{
		ID:   uuid.NewString(),
		Name: name,
	}

	if err := s.repo.Save(user); err != nil {
		return User{}, err
	}

	return user, nil
}
