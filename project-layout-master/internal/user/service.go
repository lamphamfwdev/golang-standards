package user

import (
	"database/sql"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(req RegisterRequest) (*User, error) {
	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	if req.Phone == "" {
		return nil, errors.New("phone is required")
	}

	_, err := s.repo.FindByEmail(req.Email)

	if err == nil {
		return nil, errors.New("email already exists")
	}

	if err != sql.ErrNoRows {
		return nil, err
	}

	_, err = s.repo.FindByPhone(req.Phone)
	if err == nil {
		return nil, errors.New("phone already exists")
	}

	if err != sql.ErrNoRows {
		return nil, err
	}

	return s.repo.Create(req)
}

func (s *Service) ListUser() ([]*User, error) {
	return s.repo.List()
}
