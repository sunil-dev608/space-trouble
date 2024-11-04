package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/sunil-dev608/space-trouble/config"
	"github.com/sunil-dev608/space-trouble/internal/pkg/logger"
	"github.com/sunil-dev608/space-trouble/internal/server"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log := logger.New()
	defer log.Sync()

	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file:", zap.Error(err))
	}

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("Error loading .env file:", zap.Error(err))
	}

	err = config.LoadLaunchpads()
	if err != nil {
		log.Fatal("Error Loading launchpads:", zap.Error(err))
	}
	db, err := gorm.Open(postgres.Open(cfg.DBDsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to DB:", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB from gorm.DB:", zap.Error(err))
	}
	defer sqlDB.Close()

	srv := server.New(cfg, db, log)

	// Start server
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal("server startup failed", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown failed", zap.Error(err))
	}
}
