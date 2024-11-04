package config

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	DBDsn                      string
	ServerAddress              string
	CompetitorLaunchesAPIURL   string
	CompetitorLaunchpadsAPIURL string
}

var (
	once sync.Once
	cfg  *Config
)

// GetConfig returns the configuration
func GetConfig() (*Config, error) {

	var err error
	once.Do(func() {
		cfg = &Config{}
		cfg.DBDsn = os.Getenv("DB_DSN")
		if cfg.DBDsn == "" {
			cfg = nil
			err = fmt.Errorf("DB_DSN is not set")
		}
		cfg.ServerAddress = os.Getenv("SERVER_ADDRESS")
		if cfg.ServerAddress == "" {
			cfg.ServerAddress = ":8080"
		}

		cfg.CompetitorLaunchesAPIURL = os.Getenv("COMPETITOR_LAUNCHES_API_URL")
		cfg.CompetitorLaunchpadsAPIURL = os.Getenv("COMPETITOR_LAUNCHPADS_API_URL")
	})

	return cfg, err
}
