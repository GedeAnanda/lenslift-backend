package dto

type ExerciseRequest struct {
	ExerciseName string `json:"exercise_name" binding:"required"`
	TargetSets int `json:"target_sets" binding:"required,gt=0"`
	TargetReps int `json:"target_reps" binding:"required,gt=0"`
	Notes string `json:"notes"`
	OrderIndex int `json:"order_index" binding:"required,gt=0"`
}

type WorkoutTemplateRequest struct {
	Name string`json:"name" binding:"required"`
	Description string `json:"description"`
	Exercises []ExerciseRequest `json:"exercises" binding:"required,min=1"`
}

type ExerciseResponse struct {
	ID string `json:"id"`
	ExerciseName string `json:"exercise_name"`
	TargetSets int`json:"target_sets"`
	TargetReps int `json:"target_reps"`
	Notes string `json:"notes"`
	OrderIndex int`json:"order_index"`
}

type WorkoutTemplateResponse struct {
	ID string`json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Exercises []ExerciseResponse `json:"exercises"`
}

type WorkoutTemplateListResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}