package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/model"
	"github.com/nanda/lenslift-backend/internal/repository"
	"gorm.io/gorm"
)

type ScheduleService struct {
	scheduleRepo *repository.ScheduleRepository
	workoutRepo  *repository.WorkoutRepository
}

func NewScheduleService() *ScheduleService {
	return &ScheduleService{
		scheduleRepo: repository.NewScheduleRepository(),
		workoutRepo:  repository.NewWorkoutRepository(),
	}
}

func (s *ScheduleService) SetSchedule(userID string, req dto.ScheduleRequest) (*dto.ScheduleResponse, error) {
	template, err := s.workoutRepo.FindByID(req.TemplateID)
	if err != nil {
		return nil, errors.New("template tidak ditemukan")
	}

	if template.UserID.String() != userID {
		return nil, errors.New("akses ditolak")
	}

	templateUUID, _ := uuid.Parse(req.TemplateID)
	userUUID, _ := uuid.Parse(userID)

	existing, err := s.scheduleRepo.FindByUserIDAndDay(userID, req.DayOfWeek)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existing != nil {
		existing.TemplateID = templateUUID
		if err := s.scheduleRepo.Update(existing); err != nil {
			return nil, err
		}
		return &dto.ScheduleResponse{
			ID:        existing.ID.String(),
			DayOfWeek: existing.DayOfWeek,
			Template: dto.WorkoutTemplateListResponse{
				ID:          template.ID.String(),
				Name:        template.Name,
				Description: template.Description,
			},
		}, nil
	}

	schedule := &model.GymSchedule{
		UserID:     userUUID,
		TemplateID: templateUUID,
		DayOfWeek:  req.DayOfWeek,
		IsActive:   true,
	}
	if err := s.scheduleRepo.Create(schedule); err != nil {
		return nil, err
	}

	return &dto.ScheduleResponse{
		ID:        schedule.ID.String(),
		DayOfWeek: schedule.DayOfWeek,
		Template: dto.WorkoutTemplateListResponse{
			ID:          template.ID.String(),
			Name:        template.Name,
			Description: template.Description,
		},
	}, nil
}

func (s *ScheduleService) GetSchedules(userID string) ([]dto.ScheduleResponse, error) {
	schedules, err := s.scheduleRepo.FindAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	var resp []dto.ScheduleResponse
	for _, sc := range schedules {
		resp = append(resp, dto.ScheduleResponse{
			ID:        sc.ID.String(),
			DayOfWeek: sc.DayOfWeek,
			Template: dto.WorkoutTemplateListResponse{
				ID:          sc.Template.ID.String(),
				Name:        sc.Template.Name,
				Description: sc.Template.Description,
			},
		})
	}
	return resp, nil
}

func (s *ScheduleService) DeleteSchedule(userID string, day string) error {
	existing, err := s.scheduleRepo.FindByUserIDAndDay(userID, day)
	if err != nil {
		return errors.New("jadwal tidak ditemukan")
	}

	if existing.UserID.String() != userID {
		return errors.New("akses ditolak")
	}

	return s.scheduleRepo.DeleteByDay(userID, day)
}