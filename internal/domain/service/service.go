package service

import (
	"blog-api-clean-architecture/internal/domain/models"
)

type IUserService interface {
	GetUser(id string) (*models.User, error)
	SetUser(user *models.User) error
}
