package service

import (
	"encoding/json"
	"fmt"

	"blog-api-clean-architecture/internal/domain/models"
	"blog-api-clean-architecture/internal/domain/repository"
)

type UserService struct {
	DB repository.IDBRepository
}

func NewUserService(dbRepo repository.IDBRepository) IUserService {
	return UserService{dbRepo}
}

func (s UserService) GetUser(id string) (*models.User, error) {
	value, ok := s.DB.Get(id)
	if !ok {
		return nil, nil
	}

	user, ok := value.(models.User)
	if !ok {
		return nil, fmt.Errorf("failed to get user")
	}

	return &user, nil
}

func (s UserService) SetUser(user *models.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data")
	}
	s.DB.Set("1", string(bytes))

	return nil
}
