package dto

type ScheduleRequest struct {
	DayOfWeek  string `json:"day_of_week" binding:"required,oneof=monday tuesday wednesday thursday friday saturday sunday"`
	TemplateID string `json:"template_id" binding:"required"`
}

type ScheduleResponse struct {
	ID string`json:"id"`
	DayOfWeek string `json:"day_of_week"`
	Template WorkoutTemplateListResponse`json:"template"`
}