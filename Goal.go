package googleanalytics

import (
	"fmt"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	g_types "github.com/leapforce-libraries/go_googleanalytics/types"
	go_http "github.com/leapforce-libraries/go_http"
)

type GoalResponse struct {
	Kind         string `json:"kind"`
	Username     string `json:"username"`
	TotalResults int64  `json:"totalResults"`
	StartIndex   int64  `json:"startIndex"`
	ItemsPerPage int64  `json:"itemsPerPage"`
	PreviousLink string `json:"previousLink"`
	NextLink     string `json:"nextLink"`
	Items        []Goal `json:"items"`
}

type Goal struct {
	ID                    string                 `json:"id"`
	Kind                  string                 `json:"kind"`
	SelfLink              string                 `json:"selfLink"`
	AccountID             string                 `json:"accountId"`
	WebPropertyID         string                 `json:"webPropertyId"`
	InternalWebPropertyID string                 `json:"internalWebPropertyId"`
	ProfileID             string                 `json:"profileId"`
	Name                  string                 `json:"name"`
	Value                 float64                `json:"value"`
	Active                bool                   `json:"active"`
	Type                  string                 `json:"type"`
	Created               g_types.DateTimeString `json:"created"`
	Updated               g_types.DateTimeString `json:"updated"`
	ParentLink            struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"parentLink"`
}

type ListGoalsConfig struct {
	AccountID     string
	WebPropertyID string
	ViewID        string
	MaxResults    *int64
	StartIndex    *int64
}

func (service *Service) ListGoals(config *ListGoalsConfig) (*[]Goal, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("ListGoalsConfig must not be nil")
	}
	values := url.Values{}

	if config != nil {
		if config.MaxResults != nil {
			values.Set("max-results", fmt.Sprintf("%v", *config.MaxResults))
		}
		if config.StartIndex != nil {
			values.Set("start-index", fmt.Sprintf("%v", *config.StartIndex))
		}
	}

	url := service.urlAnalytics(fmt.Sprintf("management/accounts/%s/webproperties/%s/profiles/%s/goals?%s", config.AccountID, config.WebPropertyID, config.ViewID, values.Encode()))

	goals := []Goal{}

	for {
		response := GoalResponse{}

		requestConfig := go_http.RequestConfig{
			URL:           url,
			ResponseModel: &response,
		}

		_, _, e := service.googleService.Get(&requestConfig)
		if e != nil {
			return nil, e
		}

		if len(response.Items) == 0 {
			break
		}

		goals = append(goals, response.Items...)

		if config.StartIndex != nil {
			break
		}

		if response.NextLink == "" {
			break
		}

		url = response.NextLink
	}

	return &goals, nil
}
