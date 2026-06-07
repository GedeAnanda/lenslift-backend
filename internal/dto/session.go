package dto

import "time"

type StartSessionRequest struct {
	TemplateID string `json:"template_id"`
}

type LogSetRequest struct {
	ExerciseName string`json:"exercise_name" binding:"required"`
	SetNumber int`json:"set_number" binding:"required,gt=0"`
	ActualReps int`json:"actual_reps" binding:"required,gt=0"`
	ActualWeightKg float64 `json:"actual_weight_kg"`
}

type SessionResponse struct {
	ID string `json:"id"`
	TemplateID string `json:"template_id,omitempty"`
	StartedAt time.Time`json:"started_at"`
	EndedAt *time.Time`json:"ended_at"`
	Notes string`json:"notes"`
}

type SessionLogResponse struct {
	ID string `json:"id"`
	ExerciseName string`json:"exercise_name"`
	SetNumber int`json:"set_number"`
	ActualReps int `json:"actual_reps"`
	ActualWeightKg float64`json:"actual_weight_kg"`
	LoggedAt time.Time `json:"logged_at"`
}

type SessionDetailResponse struct {
	ID string`json:"id"`
	TemplateID string`json:"template_id,omitempty"`
	StartedAt time.Time`json:"started_at"`
	EndedAt *time.Time `json:"ended_at"`
	DurationMinutes int `json:"duration_minutes,omitempty"`
	TotalSets int `json:"total_sets"`
	TotalVolumeKg float64 `json:"total_volume_kg"`
	Logs []SessionLogResponse `json:"logs"`
}