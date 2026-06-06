package dto

type UpdateProfileRequest struct {
	FullName string`json:"full_name"`
	WeightKg float64 `json:"weight_kg" binding:"required,gt=0"`
	HeightCm float64`json:"height_cm" binding:"required,gt=0"`
	Age int `json:"age" binding:"required,gt=0"`
	Gender string `json:"gender" binding:"required,oneof=male female"`
	Goal string`json:"goal" binding:"required,oneof=cut bulk maintain"`
}

type ProfileResponse struct {
	ID string`json:"id"`
	FullName string`json:"full_name"`
	WeightKg float64`json:"weight_kg"`
	HeightCm float64`json:"height_cm"`
	Age int`json:"age"`
	Gender string `json:"gender"`
	Goal string`json:"goal"`
	TargetCalories int`json:"target_calories"`
	TargetProteinG int`json:"target_protein_g"`
}