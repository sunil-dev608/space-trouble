package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/sunil-dev608/space-trouble/internal/competitors"
)

type Config struct {
	DBDsn                      string
	ServerAddress              string
	CompetitorLaunchesAPIURL   string
	CompetitorLaunchpadsAPIURL string
	Destinations               map[string]interface{}
	Launchpads                 map[string]string
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
		} else {
			cfg.ServerAddress = os.Getenv("SERVER_ADDRESS")
			if cfg.ServerAddress == "" {
				cfg.ServerAddress = ":8080"
			}

			cfg.Destinations = make(map[string]interface{})
			destinations := strings.Split(os.Getenv("DESTINATIONS"), ",")
			for _, destination := range destinations {
				cfg.Destinations[destination] = nil
			}
			cfg.CompetitorLaunchesAPIURL = os.Getenv("COMPETITOR_LAUNCHES_API_URL")
			cfg.CompetitorLaunchpadsAPIURL = os.Getenv("COMPETITOR_LAUNCHPADS_API_URL")
		}

	})

	return cfg, err
}

func LoadLaunchpads() error {
	competitorLaunchpadProvider := competitors.NewCompetitorLaunchpadsProvier(cfg.CompetitorLaunchpadsAPIURL)

	launchpads, err := competitorLaunchpadProvider.FetchLaunchpads()
	if err != nil {
		return err
	}
	cfg.Launchpads = launchpads
	return nil
}
