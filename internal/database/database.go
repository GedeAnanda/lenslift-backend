package database

import (
	"fmt"
	"log"
	"os"

	"github.com/nanda/lenslift-backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("gagal konek ke database:", err)
	}

	DB = db
	log.Println("database terhubung")
	migrate(db)

}


func migrate(db *gorm.DB) {
	log.Println("menjalankan migration...")

	err := db.AutoMigrate(
		&model.User{},
		&model.Profile{},
		&model.WorkoutTemplate{},
		&model.TemplateExercise{},
		&model.GymSchedule{},
		&model.WorkoutSession{},
		&model.SessionLog{},
		&model.FoodLog{},
		&model.BodyWeight{},
	)
	if err != nil {
		log.Fatal("migration gagal:", err)
	}

	db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_gym_schedules_user_day ON gym_schedules (user_id, day_of_week) WHERE is_active = true`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_food_logs_user_date ON food_logs (user_id, log_date)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_body_weights_user_date ON body_weights (user_id, measured_date)`)

	log.Println("migration selesai")
}
