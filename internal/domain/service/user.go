package service

import (
	"encoding/json"
	"fmt"

	"blog-api-clean-architecture/internal/domain/models"
	"blog-api-clean-architecture/internal/domain/repository"

	"github.com/google/uuid"
)

type UserService struct {
	DBRepo repository.IDBRepository
}

func NewUserService(dbRepo repository.IDBRepository) IUserService {
	return UserService{dbRepo}
}

// IUserService : ドメインモデルuserの操作を担当する。どのようなDBかの知識は持たない。
type IUserService interface {
	FindByID(id string) (*models.User, error)
	Save(user *models.User) error
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

func (s UserService) Save(user *models.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data")
	}

	id := uuid.NewString()
	if err := s.DBRepo.Set(id, string(bytes)); err != nil {
		return fmt.Errorf("failed to set user data: %w", err)
	}

	return nil
}
