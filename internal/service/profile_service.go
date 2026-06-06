package service

import (
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/repository"
)

type ProfileService struct {
	profileRepo *repository.ProfileRepository
}

func NewProfileService() *ProfileService {
	return &ProfileService{
		profileRepo: repository.NewProfileRepository(),
	}
}

func (s *ProfileService) GetProfile(userID string) (*dto.ProfileResponse, error) {
	profile, err := s.profileRepo.FindByUserID(userID)
	if err != nil {
		return nil,err
	}
	return &dto.ProfileResponse{
		ID:profile.ID.String(), 
		FullName: profile.FullName,
		WeightKg: profile.WeightKg,
		HeightCm: profile.HeightCm,
		Age: profile.Age,
		Gender: profile.Gender,
		Goal: profile.Goal,
		TargetCalories: profile.TargetCalories,
		TargetProteinG: profile.TargetProteinG,
	}, nil
}

func (s *ProfileService) UpdateProfile(userID string, req dto.UpdateProfileRequest) (*dto.ProfileResponse, error) {
	profile, err := s.profileRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	profile.FullName = req.FullName
	profile.WeightKg = req.WeightKg
	profile.HeightCm = req.HeightCm
	profile.Age = req.Age
	profile.Gender = req.Gender
	profile.Goal = req.Goal
	profile.TargetCalories = calculateCalories(req)
	profile.TargetProteinG = calculateProtein(req.WeightKg)

	if err:= s.profileRepo.Update(profile); err != nil {
		return nil, err
	}

	return &dto.ProfileResponse{
		ID:profile.ID.String(), 
		FullName: profile.FullName,
		WeightKg: profile.WeightKg,
		HeightCm: profile.HeightCm,
		Age: profile.Age,
		Gender: profile.Gender,
		Goal: profile.Goal,
		TargetCalories: profile.TargetCalories,
		TargetProteinG: profile.TargetProteinG,
	}, nil
}	


func calculateCalories(req dto.UpdateProfileRequest) int {
	var bmr float64 
	if req.Gender == "male" {
		bmr = 88.36 + (13.4 * req.WeightKg) + (4.8 * req.HeightCm) - (5.7 * float64(req.Age))
	} else {
		bmr = 447.6 + (9.25 * req.WeightKg) + (3.1 * req.HeightCm) - (4.33 * float64(req.Age))
	}

	tdee := bmr * 1.55
	switch req.Goal {
	case "cut":
		return int(tdee - 300)
	case "bulk":
		return int(tdee + 300)
	default:
		return int(tdee)
	}
}

func calculateProtein(weightKg float64) int {
	return int(weightKg * 2)
}