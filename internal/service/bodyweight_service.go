package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/model"
	"github.com/nanda/lenslift-backend/internal/repository"
)

type BodyWeightService struct {
	bodyWeightRepo *repository.BodyWeightRepository
}

func NewBodyWeightService() *BodyWeightService {
	return &BodyWeightService{
		bodyWeightRepo: repository.NewBodyWeightRepository(),
	}
}

func (s *BodyWeightService) LogWeight(userID string, req dto.BodyWeightRequest) (*dto.BodyWeightResponse, error) {
	userUUID, _ := uuid.Parse(userID)

	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)

	bw := &model.BodyWeight{
		UserID: userUUID,
		WeightKg: req.WeightKg,
		MeasuredDate: today,
		Notes: req.Notes,
	}

	if err := s.bodyWeightRepo.Create(bw); err != nil {
		return nil, err
	}

	return &dto.BodyWeightResponse{
		ID: bw.ID.String(),
		WeightKg: bw.WeightKg,
		MeasuredDate: bw.MeasuredDate,
		Notes: bw.Notes,
		CreatedAt: bw.CreatedAt,
	}, nil
}

func (s *BodyWeightService) GetHistory(userID string) ([]dto.BodyWeightResponse, error) {
	weights, err := s.bodyWeightRepo.FindAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	var resp []dto.BodyWeightResponse
	for _, w := range weights {
		resp = append(resp, dto.BodyWeightResponse{
			ID: w.ID.String(),
			WeightKg: w.WeightKg,
			MeasuredDate: w.MeasuredDate,
			Notes: w.Notes,
			CreatedAt: w.CreatedAt,
		})
	}

	if resp == nil {
		resp = []dto.BodyWeightResponse{}
	}

	return resp, nil
}

func (s *BodyWeightService) GetLatest(userID string) (*dto.BodyWeightResponse, error) {
	weight, err := s.bodyWeightRepo.FindLatestByUserID(userID)
	if err != nil {
		return nil, errors.New("belum ada data berat badan")
	}

	return &dto.BodyWeightResponse{
		ID: weight.ID.String(),
		WeightKg: weight.WeightKg,
		MeasuredDate: weight.MeasuredDate,
		Notes: weight.Notes,
		CreatedAt: weight.CreatedAt,
	}, nil
}

func (s *BodyWeightService) DeleteWeight(userID string, weightID string) error {
	weights, err := s.bodyWeightRepo.FindAllByUserID(userID)
	if err != nil {
		return err
	}

	found := false
	for _, w := range weights {
		if w.ID.String() == weightID {
			found = true
			if w.UserID.String() != userID {
				return errors.New("akses ditolak")
			}
			break
		}
	}

	if !found {
		return errors.New("data tidak ditemukan")
	}

	return s.bodyWeightRepo.Delete(weightID)
}