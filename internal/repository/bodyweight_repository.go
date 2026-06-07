package repository

import (
	"github.com/nanda/lenslift-backend/internal/database"
	"github.com/nanda/lenslift-backend/internal/model"
)

type BodyWeightRepository struct{}

func NewBodyWeightRepository() *BodyWeightRepository {
	return &BodyWeightRepository{}
}

func (r *BodyWeightRepository) Create(bw *model.BodyWeight) error {
	result := database.DB.Create(bw)
	return result.Error
}

func (r *BodyWeightRepository) FindAllByUserID(userID string) ([]model.BodyWeight, error) {
	var weights []model.BodyWeight
	result := database.DB.Where("user_id = ?", userID).Order("measured_date desc").Find(&weights)
	if result.Error != nil {
		return nil, result.Error
	}
	return weights, nil
}

func (r *BodyWeightRepository) FindLatestByUserID(userID string) (*model.BodyWeight, error) {
	var weight model.BodyWeight
	result := database.DB.Where("user_id = ?", userID).Order("measured_date desc").First(&weight)
	if result.Error != nil {
		return nil, result.Error
	}
	return &weight, nil
}

func (r *BodyWeightRepository) Delete(id string) error {
	result := database.DB.Where("id = ?", id).Delete(&model.BodyWeight{})
	return result.Error
}