package competitors

import (
	"context"
	"net/http"

	"github.com/sunil-dev608/space-trouble/internal/pkg/apicalls"
)

type LaunchpadResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type CompetitorLaunchpadsProvier interface {
	FetchLaunchpads() (map[string]string, error)
}

type competitorLaunchpadsProvier struct {
	APIURL       string
	Client       *http.Client
	lauchpadsMap map[string]string
}

func NewCompetitorLaunchpadsProvier(apiURL string) CompetitorLaunchpadsProvier {
	return &competitorLaunchpadsProvier{
		APIURL:       apiURL,
		Client:       &http.Client{},
		lauchpadsMap: make(map[string]string),
	}
}

const (
	LaunchpadActive            = "active"
	LaunchpadRetired           = "retired"
	LaunchpadUnderConstruction = "under construction"
)

func (p *competitorLaunchpadsProvier) FetchLaunchpads() (map[string]string, error) {

	var launchpads []LaunchpadResponse
	err := apicalls.APICall(context.Background(), "GET", p.Client, []byte{}, p.APIURL, &launchpads)
	if err != nil {
		return nil, err
	}

	var lauchPadsMap = make(map[string]string)

	for _, launchpad := range launchpads {
		lauchPadsMap[launchpad.ID] = launchpad.Status
	}

	// config.SetLaunchpads(lauchPadsMap)
	return lauchPadsMap, nil
}
