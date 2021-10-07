package googleanalytics

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	g_types "github.com/leapforce-libraries/go_googleanalytics/types"
	go_http "github.com/leapforce-libraries/go_http"
)

type ViewResponse struct {
	Kind         string `json:"kind"`
	Username     string `json:"username"`
	TotalResults int64  `json:"totalResults"`
	StartIndex   int64  `json:"startIndex"`
	ItemsPerPage int64  `json:"itemsPerPage"`
	PreviousLink string `json:"previousLink"`
	NextLink     string `json:"nextLink"`
	Items        []View `json:"items"`
}

type View struct {
	ID                                string `json:"id"`
	Kind                              string `json:"kind"`
	SelfLink                          string `json:"selfLink"`
	AccountID                         string `json:"accountId"`
	WebPropertyID                     string `json:"webPropertyId"`
	InternalWebPropertyID             string `json:"internalWebPropertyId"`
	Name                              string `json:"name"`
	Currency                          string `json:"currency"`
	Timezone                          string `json:"timezone"`
	WebsiteURL                        string `json:"websiteUrl"`
	DefaultPage                       string `json:"defaultPage"`
	ExcludeQueryParameters            string `json:"excludeQueryParameters"`
	SiteSearchQueryParameters         string `json:"siteSearchQueryParameters"`
	StripSiteSearchQueryParameters    bool   `json:"stripSiteSearchQueryParameters"`
	SiteSearchCategoryParameters      string `json:"siteSearchCategoryParameters"`
	StripSiteSearchCategoryParameters bool   `json:"stripSiteSearchCategoryParameters"`
	Type                              string `json:"type"`
	Permissions                       struct {
		Effective []string `json:"effective"`
	} `json:"permissions"`
	Created                   g_types.DateTimeString `json:"created"`
	Updated                   g_types.DateTimeString `json:"updated"`
	ECommerceTracking         bool                   `json:"eCommerceTracking"`
	EnhancedECommerceTracking bool                   `json:"enhancedECommerceTracking"`
	BotFilteringEnabled       bool                   `json:"botFilteringEnabled"`
	Starred                   bool                   `json:"starred"`
	ParentLink                struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"parentLink"`
	ChildLink struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"childLink"`
}

type ListViewsConfig struct {
	AccountID     string
	WebPropertyID string
	MaxResults    *int64
	StartIndex    *int64
}

func (service *Service) ListViews(config *ListViewsConfig) (*[]View, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("ListViewsConfig must not be nil")
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

	url := service.urlAnalytics(fmt.Sprintf("management/accounts/%s/webproperties/%s/profiles?%s", config.AccountID, config.WebPropertyID, values.Encode()))

	views := []View{}

	for {
		response := ViewResponse{}

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

		views = append(views, response.Items...)

		if config.StartIndex != nil {
			break
		}

		if response.NextLink == "" {
			break
		}

		url = response.NextLink
	}

	return &views, nil
}
