package model

import (
	"time"

	"github.com/google/uuid"
)

type WorkoutTemplate struct { 
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID `gorm:"type:uuid;index;not null"`
	Name string `gorm:"size:100;not null"`
	Description string `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	User User `gorm:"foreignKey: UserId"`
	
	Exercises []TemplateExercise `gorm:"foreignKey:TemplateID"`
	Sessions []WorkoutSession`gorm:"foreignKey:TemplateID"`
	Schedules []GymSchedule`gorm:"foreignKey:TemplateID"`
	
}

type TemplateExercise struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TemplateID uuid.UUID `gorm:"type:uuid;index;not null"`
	ExerciseName string `gorm:"size:100;not null"`
	TargetSets int`gorm:"not null"`
	TargetReps int `gorm:"not null"`
	Notes string `gorm:"type:text"`
	OrderIndex int `gorm:"not null;default:0"`

	Template WorkoutTemplate `gorm:"foreignKey:TemplateID"`
}

type GymSchedule struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID `gorm:"type:uuid;index;not null"`
	TemplateID uuid.UUID `gorm:"type:uuid;index;not null"`
	DayOfWeek string `gorm:"size:10;not null;check:day_of_week IN ('monday','tuesday','wednesday','thursday','friday','saturday','sunday')"`
	IsActive bool `gorm:"default:true"`

	User User`gorm:"foreignKey:UserID"`
	Template WorkoutTemplate`gorm:"foreignKey:TemplateID"`
}

type WorkoutSession struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID `gorm:"type:uuid;index;not null"`
	TemplateID *uuid.UUID `gorm:"type:uuid;index"`
	StartedAt time.Time`gorm:"autoCreateTime"`
	EndedAt *time.Time
	Notes string `gorm:"type:text"`

	User User`gorm:"foreignKey:UserID"`
	Template *WorkoutTemplate `gorm:"foreignKey:TemplateID"`
	Logs []SessionLog`gorm:"foreignKey:SessionID"`
}

type SessionLog struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SessionID uuid.UUID `gorm:"type:uuid;index;not null"`
	ExerciseID *uuid.UUID `gorm:"type:uuid;index"`
	ExerciseName string `gorm:"size:100;not null"`
	SetNumber int `gorm:"not null"`
	ActualReps int `gorm:"not null"`
	ActualWeightKg float64 `gorm:"type:decimal(6,2)"`
	LoggedAt time.Time`gorm:"autoCreateTime"`
	
	Session WorkoutSession `gorm:"foreignKey:SessionID"`
	Exercise *TemplateExercise `gorm:"foreignKey:ExerciseID"`
}