package repository

import (
	"github.com/nanda/lenslift-backend/internal/database"
	"github.com/nanda/lenslift-backend/internal/model"
)

type ScheduleRepository struct{}

func NewScheduleRepository() *ScheduleRepository {
	return &ScheduleRepository{}
}

func (r *ScheduleRepository) FindByUserIDAndDay(userID string, day string) (*model.GymSchedule, error) {
	var schedule model.GymSchedule
	result := database.DB.Where("user_id = ? AND day_of_week = ? AND is_active = true", userID, day).First(&schedule)
	if result.Error != nil {
		return nil, result.Error
	}
	return &schedule, nil
}

func (r *ScheduleRepository) FindAllByUserID(userID string) ([]model.GymSchedule, error) {
	var schedules []model.GymSchedule
	result := database.DB.Preload("Template").Where("user_id = ? AND is_active = true", userID).Find(&schedules)
	if result.Error != nil {
		return nil, result.Error
	}
	return schedules, nil
}

func (r *ScheduleRepository) Create(schedule *model.GymSchedule) error {
	result := database.DB.Create(schedule)
	return result.Error
}

func (r *ScheduleRepository) Update(schedule *model.GymSchedule) error {
	result := database.DB.Save(schedule)
	return result.Error
}

func (r *ScheduleRepository) DeleteByDay(userID string, day string) error {
	result := database.DB.Where("user_id = ? AND day_of_week = ?", userID, day).Delete(&model.GymSchedule{})
	return result.Error
}