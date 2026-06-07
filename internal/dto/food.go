package dto

import "time"

type FoodLogRequest struct {
	FoodName string `json:"food_name" binding:"required"`
	Calories float64 `json:"calories" binding:"required,gt=0"`
	ProteinG float64 `json:"protein_g" binding:"required,gt=0"`
	CarbsG float64 `json:"carbs_g"`
	FatG float64 `json:"fat_g"`
	LogDate string `json:"log_date"`
}

type FoodLogResponse struct {
	ID string `json:"id"`
	FoodName string `json:"food_name"`
	Calories float64 `json:"calories"`
	ProteinG float64 `json:"protein_g"`
	CarbsG float64 `json:"carbs_g"`
	FatG float64 `json:"fat_g"`
	Source string `json:"source"`
	LogDate time.Time `json:"log_date"`
	CreatedAt time.Time `json:"created_at"`
}

type DailySummary struct {
	TotalCalories float64 `json:"total_calories"`
	TargetCalories int `json:"target_calories"`
	TotalProteinG float64 `json:"total_protein_g"`
	TargetProteinG int `json:"target_protein_g"`
	TotalCarbsG float64 `json:"total_carbs_g"`
	TotalFatG float64 `json:"total_fat_g"`
}

type FoodLogWithSummary struct {
	FoodLog FoodLogResponse `json:"food_log"`
	DailySummary DailySummary `json:"daily_summary"`
}

type DailyFoodLogs struct {
	Logs []FoodLogResponse `json:"logs"`
	DailySummary DailySummary `json:"daily_summary"`
}

type AIAnalyzeResponse struct {
	FoodName string `json:"food_name"`
	Calories float64 `json:"calories"`
	ProteinG float64 `json:"protein_g"`
	CarbsG float64 `json:"carbs_g"`
	FatG float64 `json:"fat_g"`
}