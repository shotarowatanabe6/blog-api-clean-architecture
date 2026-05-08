package service

import (
	"encoding/json"
	"fmt"
	"time"

	"blog-api-clean-architecture/internal/domain/models"
	"blog-api-clean-architecture/internal/domain/repository"

	"github.com/google/uuid"
)

type UserService struct {
	DBRepo repository.IDBRepository
}

func NewUserService(dbRepo repository.IDBRepository) UserService {
	return UserService{dbRepo}
}

func (s UserService) FindByID(id string) (*models.User, error) {
	value, err := s.DBRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get value: %w", err)
	}
	if value == nil {
		return nil, nil
	}

	var user models.User
	if err = json.Unmarshal([]byte(*value), &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &user, nil
}

func (s UserService) Save(user *models.User) (*models.User, error) {
	id := uuid.NewString()
	u := &models.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: float64(time.Now().UnixMilli()),
	}

	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}
	if err := s.DBRepo.Set(id, string(bytes)); err != nil {
		return nil, fmt.Errorf("failed to set user data: %w", err)
	}

	return u, nil
}
