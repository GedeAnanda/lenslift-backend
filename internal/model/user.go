package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt time.Time`gorm:"autoCreateTime"`
	UpdatedAt time.Time`gorm:"autoUpdateTime"`

	Profile *Profile
	WorkoutTemplates []WorkoutTemplate
	WorkoutSessions []WorkoutSession
	FoodLogs []FoodLog
	BodyWeights []BodyWeight
	GymSchedules []GymSchedule
}

type Profile struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`
	FullName string`gorm:"size:100"`
	WeightKg float64`gorm:"type:decimal(5,2)"`
	HeightCm float64 `gorm:"type:decimal(5,2)"`
	Age int
	Gender string `gorm:"size:10"`
	Goal string `gorm:"size:10"`
	TargetCalories int
	TargetProteinG int
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	User *User `gorm:"foreignKey:UserID"`
}