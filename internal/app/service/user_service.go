package service

import (
	"github.com/mazezen/echoframe/internal/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser handles the creation logic for User
func (this *UserService) CreateUser(input *models.User) error {
	return dto.NewUserDto().CreateUser(input)
}

// GetUser handles the retrieval logic for User
func (this *UserService) GetUser(id string) (*models.User, error) {
	return dto.NewUserDto().GetUser(id)
}

// UpdateUser handles the update logic for User
func (this *UserService) UpdateUser(input *models.User) error {
	return dto.NewUserDto().UpdateUser(input)
}

// DeleteUser handles the deletion logic for User
func (this *UserService) DeleteUser(id string) error {
	return dto.NewUserDto().DeleteUser(id)
}
