package repository

import (
	"github.com/nanda/lenslift-backend/internal/database"
	"github.com/nanda/lenslift-backend/internal/model"
)

type ProfileRepository struct{}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (r *ProfileRepository) Create (profile *model.Profile) error {
	result := database.DB.Create(profile)
	return result.Error
}

func (r *ProfileRepository) FindByUserID(userID string) (*model.Profile, error) {
	var profile model.Profile
	result := database.DB.Where("user_id = ?", userID).First(&profile)
	if result.Error != nil {
		return nil, result.Error
	}
	return &profile, nil
}

func (r *ProfileRepository) Update(profile *model.Profile) error {
	result := database.DB.Save(profile)
	return result.Error
}