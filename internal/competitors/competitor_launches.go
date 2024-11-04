package competitors

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sunil-dev608/space-trouble/internal/pkg/apicalls"
)

type Competitor struct {
	Name string `json:"name"`
}

type DateLocal struct {
	DateGTE string `json:"$gte"`
	DateLTE string `json:"$lte"`
}

type Query struct {
	Launchpad string    `json:"launchpad"`
	DateLocal DateLocal `json:"date_local"`
}

type CompetitorLaunchesQuery struct {
	Query Query `json:"query"`
}

type CompetitorLaunchesAPIResponse struct {
	Docs []interface{} `json:"docs"`
}

type CompetitorLaunchesProvier interface {
	HasCompetingFlight(ctx context.Context, launchpadID string, launchDate time.Time) (bool, error)
}

type competitorLaunchesProvier struct {
	APIURL string
	Client *http.Client
}

func NewCompetitorLaunchesProvier(apiURL string) CompetitorLaunchesProvier {
	return &competitorLaunchesProvier{
		APIURL: apiURL,
		Client: &http.Client{},
	}
}

func (p *competitorLaunchesProvier) HasCompetingFlight(ctx context.Context, launchpadID string, launchDate time.Time) (bool, error) {
	query := CompetitorLaunchesQuery{
		Query: Query{
			Launchpad: launchpadID,
			DateLocal: DateLocal{
				DateGTE: launchDate.String(),
				DateLTE: launchDate.Add(24 * time.Hour).String(),
			},
		},
	}
	jsonData, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	var apiResponse CompetitorLaunchesAPIResponse

	if err = apicalls.APICall(ctx, "POST", p.Client, jsonData, p.APIURL, &apiResponse); err != nil {
		return false, err
	}

	return len(apiResponse.Docs) > 0, nil
}
