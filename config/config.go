package config

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	DBDsn string
}

var (
	once sync.Once
	cfg  *Config
)

func GetConfig() (*Config, error) {

	var err error
	once.Do(func() {
		cfg = &Config{}
		cfg.DBDsn = os.Getenv("DB_DSN")
		if cfg.DBDsn == "" {
			cfg = nil
			err = fmt.Errorf("DB_DSN is not set")
		}
	})

	return cfg, err
}
