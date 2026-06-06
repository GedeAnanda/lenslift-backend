package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nanda/lenslift-backend/internal/database"
	"github.com/nanda/lenslift-backend/internal/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("file .env tidak ditemukan, pakai environment variable sistem")
	}

	database.Connect()

	r := handler.NewRouter()
	port := os.Getenv("APP_PORT")
	log.Println("lenslift backend jalan di port", port)
	r.Run(":" + port)
}