package config

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	DBDsn         string
	ServerAddress string
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
	})

	return cfg, err
}
