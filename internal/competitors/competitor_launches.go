package competitors

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
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

const CompetitorQueryDateTimeFormat = "2017-12-15T00:00:00.000Z"

func (p *competitorLaunchesProvier) HasCompetingFlight(ctx context.Context, launchpadID string, launchDate time.Time) (bool, error) {
	dGTE := launchDate.String()
	dLTE := launchDate.Add(24 * time.Hour).String()
	query := CompetitorLaunchesQuery{
		Query: Query{
			Launchpad: launchpadID,
			DateLocal: DateLocal{
				DateGTE: dGTE,
				DateLTE: dLTE,
			},
		},
	}
	jsonData, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.APIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := p.Client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var apiResponse CompetitorLaunchesAPIResponse

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return false, err
	}

	return len(apiResponse.Docs) > 0, nil
}
