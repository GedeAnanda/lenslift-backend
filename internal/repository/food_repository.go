package repository

import (
	"time"

	"github.com/nanda/lenslift-backend/internal/database"
	"github.com/nanda/lenslift-backend/internal/model"
)

type FoodRepository struct{}

func NewFoodRepository() *FoodRepository {
	return &FoodRepository{}
}

func (r *FoodRepository) Create(food *model.FoodLog) error {
	result := database.DB.Create(food)
	return result.Error
}

func (r *FoodRepository) FindByUserIDAndDate(userID string, date time.Time) ([]model.FoodLog, error) {
	var logs []model.FoodLog
	result := database.DB.Where("user_id = ? AND log_date = ?", userID, date.Format("2006-01-02")).
		Order("created_at asc").
		Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

func (r *FoodRepository) FindByID(id string) (*model.FoodLog, error) {
	var food model.FoodLog
	result := database.DB.Where("id = ?", id).First(&food)
	if result.Error != nil {
		return nil, result.Error
	}
	return &food, nil
}

func (r *FoodRepository) Delete(id string) error {
	result := database.DB.Where("id = ?", id).Delete(&model.FoodLog{})
	return result.Error
}