package repository

import (
	"github.com/nanda/lenslift-backend/internal/database"
	"github.com/nanda/lenslift-backend/internal/model"
)

type SessionRepository struct{}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{}
}

func (r *SessionRepository) Create(session *model.WorkoutSession) error {
	result := database.DB.Create(session)
	return result.Error
}

func (r *SessionRepository) FindByID(id string) (*model.WorkoutSession, error) {
	var session model.WorkoutSession
	result := database.DB.Preload("Logs").Where("id = ?", id).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

func (r *SessionRepository) FindAllByUserID(userID string) ([]model.WorkoutSession, error) {
	var sessions []model.WorkoutSession
	result := database.DB.Where("user_id = ?", userID).Order("started_at desc").Find(&sessions)
	if result.Error != nil {
		return nil, result.Error
	}
	return sessions, nil
}

func (r *SessionRepository) Update(session *model.WorkoutSession) error {
	result := database.DB.Save(session)
	return result.Error
}

func (r *SessionRepository) CreateLog(log *model.SessionLog) error {
	result := database.DB.Create(log)
	return result.Error
}