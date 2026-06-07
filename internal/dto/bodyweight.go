package dto

import "time"

type BodyWeightRequest struct {
	WeightKg float64 `json:"weight_kg" binding:"required,gt=0"`
	Notes string `json:"notes"`
}

type BodyWeightResponse struct {
	ID string `json:"id"`
	WeightKg float64 `json:"weight_kg"`
	MeasuredDate time.Time `json:"measured_date"`
	Notes string `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}