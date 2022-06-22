package googleanalytics

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	g_types "github.com/leapforce-libraries/go_googleanalytics/types"
	go_http "github.com/leapforce-libraries/go_http"
)

type PropertyResponse struct {
	Properties    []Property `json:"properties"`
	NextPageToken *string    `json:"nextPageToken"`
}

type Property struct {
	Name             string                 `json:"name"`
	CreateTime       g_types.DateTimeString `json:"createTime"`
	UpdateTime       g_types.DateTimeString `json:"updateTime"`
	Parent           string                 `json:"parent"`
	DisplayName      string                 `json:"displayName"`
	IndustryCategory IndustryCategory       `json:"industryCategory"`
	TimeZone         string                 `json:"timeZone"`
	CurrencyCode     string                 `json:"currencyCode"`
	ServiceLevel     ServiceLevel           `json:"serviceLevel"`
	DeleteTime       g_types.DateTimeString `json:"deleteTime"`
	ExpireTime       g_types.DateTimeString `json:"expireTime"`
	Account          string                 `json:"account"`
}

type IndustryCategory string

const (
	IndustryCategoryIndustryCategoryUnspecified  IndustryCategory = "INDUSTRY_CATEGORY_UNSPECIFIED"
	IndustryCategoryAutomotive                   IndustryCategory = "AUTOMOTIVE"
	IndustryCategoryBusinessAndIndustrialMarkets IndustryCategory = "BUSINESS_AND_INDUSTRIAL_MARKETS"
	IndustryCategoryFinance                      IndustryCategory = "FINANCE"
	IndustryCategoryHealthcare                   IndustryCategory = "HEALTHCARE"
	IndustryCategoryTechnology                   IndustryCategory = "TECHNOLOGY"
	IndustryCategoryTravel                       IndustryCategory = "TRAVEL"
	IndustryCategoryOther                        IndustryCategory = "OTHER"
	IndustryCategoryArtsAndEntertainment         IndustryCategory = "ARTS_AND_ENTERTAINMENT"
	IndustryCategoryBeautyAndFitness             IndustryCategory = "BEAUTY_AND_FITNESS"
	IndustryCategoryBooksAndLiterature           IndustryCategory = "BOOKS_AND_LITERATURE"
	IndustryCategoryFoodAndDrink                 IndustryCategory = "FOOD_AND_DRINK"
	IndustryCategoryGames                        IndustryCategory = "GAMES"
	IndustryCategoryHobbiesAndLeisure            IndustryCategory = "HOBBIES_AND_LEISURE"
	IndustryCategoryHomeAndGarden                IndustryCategory = "HOME_AND_GARDEN"
	IndustryCategoryInternetAndTelecom           IndustryCategory = "INTERNET_AND_TELECOM"
	IndustryCategoryLawAndGovernment             IndustryCategory = "LAW_AND_GOVERNMENT"
	IndustryCategoryNews                         IndustryCategory = "NEWS"
	IndustryCategoryOnlineCommunities            IndustryCategory = "ONLINE_COMMUNITIES"
	IndustryCategoryPeopleAndSociety             IndustryCategory = "PEOPLE_AND_SOCIETY"
	IndustryCategoryPetsAndAnimals               IndustryCategory = "PETS_AND_ANIMALS"
	IndustryCategoryRealEstate                   IndustryCategory = "REAL_ESTATE"
	IndustryCategoryReference                    IndustryCategory = "REFERENCE"
	IndustryCategoryScience                      IndustryCategory = "SCIENCE"
	IndustryCategorySports                       IndustryCategory = "SPORTS"
	IndustryCategoryJobsAndEducation             IndustryCategory = "JOBS_AND_EDUCATION"
	IndustryCategoryShopping                     IndustryCategory = "SHOPPING"
)

type ServiceLevel string

const (
	ServiceLevelUnspecified             ServiceLevel = "SERVICE_LEVEL_UNSPECIFIED"
	ServiceLevelGoogleAnalyticsStandard ServiceLevel = "GOOGLE_ANALYTICS_STANDARD"
	ServiceLevelGoogleAnalytics360      ServiceLevel = "GOOGLE_ANALYTICS_360"
)

type ListPropertiesConfig struct {
	Filter struct {
		Parent          *string
		Ancestor        *string
		FirebaseProject *string
	}
	PageSize    *int64
	PageToken   *string
	ShowDeleted *bool
}

func (service *Service) ListProperties(config *ListPropertiesConfig) (*[]Property, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("ListPropertiesConfig must not be nil")
	}

	var pageToken *string = nil

	values := url.Values{}

	if config.Filter.Parent != nil {
		values.Set("filter", fmt.Sprintf("parent%s", *config.Filter.Parent))
	}
	if config.Filter.Ancestor != nil {
		values.Set("filter", fmt.Sprintf("ancestor:%s", *config.Filter.Ancestor))
	}
	if config.Filter.FirebaseProject != nil {
		values.Set("filter", fmt.Sprintf("firebase_project:%s", *config.Filter.FirebaseProject))
	}
	if config.PageSize != nil {
		values.Set("pageSize", fmt.Sprintf("%v", *config.PageSize))
	}
	if config.PageToken != nil {
		pageToken = config.PageToken
	}
	if config.ShowDeleted != nil {
		values.Set("showDeleted", fmt.Sprintf("%v", *config.ShowDeleted))
	}

	properties := []Property{}

	for {
		if pageToken != nil {
			values.Set("pageToken", *pageToken)
		}

		response := PropertyResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.urlAdmin(fmt.Sprintf("properties?%s", values.Encode())),
			ResponseModel: &response,
		}

		_, _, e := service.googleService().HttpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		if len(response.Properties) == 0 {
			break
		}

		properties = append(properties, response.Properties...)

		if response.NextPageToken == nil {
			break
		}

		pageToken = response.NextPageToken
	}

	return &properties, nil
}
