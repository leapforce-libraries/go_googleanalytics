package googleanalytics

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	g_types "github.com/leapforce-libraries/go_googleanalytics/types"
	go_http "github.com/leapforce-libraries/go_http"
)

type AccountResponse struct {
	Kind         string    `json:"kind"`
	Username     string    `json:"username"`
	TotalResults int64     `json:"totalResults"`
	StartIndex   int64     `json:"startIndex"`
	ItemsPerPage int64     `json:"itemsPerPage"`
	PreviousLink string    `json:"previousLink"`
	NextLink     string    `json:"nextLink"`
	Items        []Account `json:"items"`
}

type Account struct {
	ID          string `json:"id"`
	Kind        string `json:"kind"`
	SelfLink    string `json:"selfLink"`
	Name        string `json:"name"`
	Permissions struct {
		Effective []string `json:"effective"`
	} `json:"permissions"`
	Created   g_types.DateTimeString `json:"created"`
	Updated   g_types.DateTimeString `json:"updated"`
	Starred   bool                   `json:"starred"`
	ChildLink struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"childLink"`
}

type ListAccountsConfig struct {
	MaxResults *int64
	StartIndex *int64
}

func (service *Service) ListAccounts(config *ListAccountsConfig) (*[]Account, *errortools.Error) {
	values := url.Values{}

	if config != nil {
		if config.MaxResults != nil {
			values.Set("max-results", fmt.Sprintf("%v", *config.MaxResults))
		}
		if config.StartIndex != nil {
			values.Set("start-index", fmt.Sprintf("%v", *config.StartIndex))
		}
	}

	url := service.urlAnalytics(fmt.Sprintf("management/accounts?%s", values.Encode()))

	accounts := []Account{}

	for {
		response := AccountResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			URL:           url,
			ResponseModel: &response,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		if len(response.Items) == 0 {
			break
		}

		accounts = append(accounts, response.Items...)

		if config != nil {
			if config.StartIndex != nil {
				break
			}
		}

		if response.NextLink == "" {
			break
		}

		url = response.NextLink
	}

	return &accounts, nil
}
