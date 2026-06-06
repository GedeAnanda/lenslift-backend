package repository

import (
	"github.com/nanda/lenslift-backend/internal/database"
	"github.com/nanda/lenslift-backend/internal/model"
)

type UserRepository struct {}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *model.User) error {
	result:= database.DB.Create(user)
	return result.Error
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindById(id string) (*model.User, error) {
	var user model.User
	result := database.DB.Preload("Profile").Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}