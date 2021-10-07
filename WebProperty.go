package googleanalytics

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	g_types "github.com/leapforce-libraries/go_googleanalytics/types"
	go_http "github.com/leapforce-libraries/go_http"
)

type WebPropertyResponse struct {
	Kind         string        `json:"kind"`
	Username     string        `json:"username"`
	TotalResults int64         `json:"totalResults"`
	StartIndex   int64         `json:"startIndex"`
	ItemsPerPage int64         `json:"itemsPerPage"`
	PreviousLink string        `json:"previousLink"`
	NextLink     string        `json:"nextLink"`
	Items        []WebProperty `json:"items"`
}

type WebProperty struct {
	ID                              string `json:"id"`
	Kind                            string `json:"kind"`
	SelfLink                        string `json:"selfLink"`
	AccountID                       string `json:"accountId"`
	InternalWebPropertyID           string `json:"internalWebPropertyId"`
	Name                            string `json:"name"`
	WebsiteURL                      string `json:"websiteUrl"`
	Level                           string `json:"level"`
	ProfileCount                    int64  `json:"profileCount"`
	IndustryVertical                string `json:"industryVertical"`
	DefaultProfileID                string `json:"defaultProfileId"`
	DataRetentionTTL                string `json:"dataRetentionTtl"`
	DataRetentionResetOnNewActivity bool   `json:"dataRetentionResetOnNewActivity"`
	Permissions                     struct {
		Effective []string `json:"effective"`
	} `json:"permissions"`
	Created    g_types.DateTimeString `json:"created"`
	Updated    g_types.DateTimeString `json:"updated"`
	Starred    bool                   `json:"starred"`
	ParentLink struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"parentLink"`
	ChildLink struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"childLink"`
}

type ListWebPropertiesConfig struct {
	AccountID  string
	MaxResults *int64
	StartIndex *int64
}

func (service *Service) ListWebProperties(config *ListWebPropertiesConfig) (*[]WebProperty, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("ListWebPropertiesConfig must not be nil")
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

	url := service.urlAnalytics(fmt.Sprintf("management/accounts/%s/webproperties?%s", config.AccountID, values.Encode()))

	webProperties := []WebProperty{}

	for {
		response := WebPropertyResponse{}

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

		webProperties = append(webProperties, response.Items...)

		if config.StartIndex != nil {
			break
		}

		if response.NextLink == "" {
			break
		}

		url = response.NextLink
	}

	return &webProperties, nil
}
