package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/model"
	"github.com/nanda/lenslift-backend/internal/repository"
)

type WorkoutService struct {
	workoutRepo *repository.WorkoutRepository
}

func NewWorkoutService() *WorkoutService {
	return &WorkoutService{
		workoutRepo: repository.NewWorkoutRepository(),
	}
}

func (s *WorkoutService) CreateTemplate(userID string, req dto.WorkoutTemplateRequest) (*dto.WorkoutTemplateResponse, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("user id tidak valid")
	}

	template := &model.WorkoutTemplate{
		UserID:      userUUID,
		Name:        req.Name,
		Description: req.Description,
	}
	if err := s.workoutRepo.CreateTemplate(template); err != nil {
		return nil, err
	}

	exercises := make([]model.TemplateExercise, len(req.Exercises))
	for i, ex := range req.Exercises {
		exercises[i] = model.TemplateExercise{
			TemplateID:   template.ID,
			ExerciseName: ex.ExerciseName,
			TargetSets:   ex.TargetSets,
			TargetReps:   ex.TargetReps,
			Notes:        ex.Notes,
			OrderIndex:   ex.OrderIndex,
		}
	}
	if err := s.workoutRepo.CreateExercises(exercises); err != nil {
		return nil, err
	}

	return buildTemplateResponse(template, exercises), nil
}

func (s *WorkoutService) GetAllTemplates(userID string) ([]dto.WorkoutTemplateListResponse, error) {
	templates, err := s.workoutRepo.FindAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	var resp []dto.WorkoutTemplateListResponse
	for _, t := range templates {
		resp = append(resp, dto.WorkoutTemplateListResponse{
			ID:          t.ID.String(),
			Name:        t.Name,
			Description: t.Description,
		})
	}
	return resp, nil
}

func (s *WorkoutService) GetTemplate(userID string, templateID string) (*dto.WorkoutTemplateResponse, error) {
	template, err := s.workoutRepo.FindByID(templateID)
	if err != nil {
		return nil, errors.New("template tidak ditemukan")
	}

	if template.UserID.String() != userID {
		return nil, errors.New("akses ditolak")
	}

	return buildTemplateResponse(template, template.Exercises), nil
}

func (s *WorkoutService) UpdateTemplate(userID string, templateID string, req dto.WorkoutTemplateRequest) (*dto.WorkoutTemplateResponse, error) {
	template, err := s.workoutRepo.FindByID(templateID)
	if err != nil {
		return nil, errors.New("template tidak ditemukan")
	}

	if template.UserID.String() != userID {
		return nil, errors.New("akses ditolak")
	}

	template.Name = req.Name
	template.Description = req.Description
	if err := s.workoutRepo.UpdateTemplate(template); err != nil {
		return nil, err
	}

	if err := s.workoutRepo.DeleteExercisesByTemplateID(templateID); err != nil {
		return nil, err
	}

	exercises := make([]model.TemplateExercise, len(req.Exercises))
	for i, ex := range req.Exercises {
		exercises[i] = model.TemplateExercise{
			TemplateID:   template.ID,
			ExerciseName: ex.ExerciseName,
			TargetSets:   ex.TargetSets,
			TargetReps:   ex.TargetReps,
			Notes:        ex.Notes,
			OrderIndex:   ex.OrderIndex,
		}
	}
	if err := s.workoutRepo.CreateExercises(exercises); err != nil {
		return nil, err
	}

	return buildTemplateResponse(template, exercises), nil
}

func (s *WorkoutService) DeleteTemplate(userID string, templateID string) error {
	template, err := s.workoutRepo.FindByID(templateID)
	if err != nil {
		return errors.New("template tidak ditemukan")
	}

	if template.UserID.String() != userID {
		return errors.New("akses ditolak")
	}

	if err := s.workoutRepo.DeleteExercisesByTemplateID(templateID); err != nil {
		return err
	}

	return s.workoutRepo.DeleteTemplate(templateID)
}

func buildTemplateResponse(template *model.WorkoutTemplate, exercises []model.TemplateExercise) *dto.WorkoutTemplateResponse {
	exerciseResponses := make([]dto.ExerciseResponse, len(exercises))
	for i, ex := range exercises {
		exerciseResponses[i] = dto.ExerciseResponse{
			ID:           ex.ID.String(),
			ExerciseName: ex.ExerciseName,
			TargetSets:   ex.TargetSets,
			TargetReps:   ex.TargetReps,
			Notes:        ex.Notes,
			OrderIndex:   ex.OrderIndex,
		}
	}

	return &dto.WorkoutTemplateResponse{
		ID:          template.ID.String(),
		Name:        template.Name,
		Description: template.Description,
		Exercises:   exerciseResponses,
	}
}