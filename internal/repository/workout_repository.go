package repository

import (
	"github.com/nanda/lenslift-backend/internal/database"
	"github.com/nanda/lenslift-backend/internal/model"
	"gorm.io/gorm"
)

type WorkoutRepository struct{}

func NewWorkoutRepository() *WorkoutRepository {
	return &WorkoutRepository{}
}

func (r *WorkoutRepository) CreateTemplate(template *model.WorkoutTemplate) error {
	result := database.DB.Create(template)
	return result.Error
}

func (r *WorkoutRepository) CreateExercises(exercises []model.TemplateExercise) error {
	result := database.DB.Create(&exercises)
	return result.Error
}

func (r *WorkoutRepository) FindAllByUserID(userID string) ([]model.WorkoutTemplate, error) {
	var templates []model.WorkoutTemplate
	result := database.DB.Where("user_id = ?", userID).Find(&templates)
	if result.Error != nil {
		return nil, result.Error
	}
	return templates, nil
}

func (r *WorkoutRepository) FindByID(id string) (*model.WorkoutTemplate, error) {
	var template model.WorkoutTemplate
	result := database.DB.Preload("Exercises", func(db *gorm.DB) *gorm.DB {
		return db.Order("order_index asc")
	}).Where("id = ?", id).First(&template)
	if result.Error != nil {
		return nil, result.Error
	}
	return &template, nil
}

func (r *WorkoutRepository) UpdateTemplate(template *model.WorkoutTemplate) error {
	result := database.DB.Save(template)
	return result.Error
}

func (r *WorkoutRepository) DeleteExercisesByTemplateID(templateID string) error {
	result := database.DB.Where("template_id = ?", templateID).Delete(&model.TemplateExercise{})
	return result.Error
}

func (r *WorkoutRepository) DeleteTemplate(id string) error {
	result := database.DB.Where("id = ?", id).Delete(&model.WorkoutTemplate{})
	return result.Error
}