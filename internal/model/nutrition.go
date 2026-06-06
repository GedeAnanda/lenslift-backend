package model

import (
	"time"

	"github.com/google/uuid"
)

type FoodLog struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID`gorm:"type:uuid;index;not null"`
	FoodName string`gorm:"size:150;not null"`
	Calories float64 `gorm:"type:decimal(8,2);not null"`
	ProteinG float64 `gorm:"type:decimal(6,2);not null"`
	CarbsG float64 `gorm:"type:decimal(6,2)"`
	FatG float64`gorm:"type:decimal(6,2)"`
	Source string `gorm:"size:20;default:'manual';check:source IN ('manual','ai_photo')"`
	LogDate time.Time `gorm:"type:date;index;not null"`
	CreatedAt time.Time`gorm:"autoCreateTime"`

	User *User `gorm:"foreignKey:UserID"`
}

type BodyWeight struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID`gorm:"type:uuid;index;not null"`
	WeightKg float64`gorm:"type:decimal(5,2);not null"`
	MeasuredDate time.Time `gorm:"type:date;index;not null"`
	Notes string`gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User *User `gorm:"foreignKey:UserID"`
}