package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sunil-dev608/space-trouble/config"
	"github.com/sunil-dev608/space-trouble/internal/pkg/repository"
	"github.com/sunil-dev608/space-trouble/internal/pkg/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	db, err := gorm.Open(postgres.Open(cfg.DBDsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	bookingRepo := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepo)
	_ = bookingService
}
