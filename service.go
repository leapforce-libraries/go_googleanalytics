package googleanalytics

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
	go_http "github.com/leapforce-libraries/go_http"
	"github.com/leapforce-libraries/go_oauth2/tokensource"
)

const (
	apiName         string = "GoogleAnalytics"
	apiURLReporting string = "https://analyticsreporting.googleapis.com/v4"
	apiURLAnalytics string = "https://www.googleapis.com/analytics/v3"
)

// Service
//
type Service struct {
	authorizationMode go_google.AuthorizationMode
	id                string
	apiKey            *string
	accessToken       *string
	httpService       *go_http.Service
	googleService     *go_google.Service
}

type ServiceConfigWithAPIKey struct {
	APIKey string
}

func NewServiceWithAPIKey(serviceConfig *ServiceConfigWithAPIKey) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.APIKey == "" {
		return nil, errortools.ErrorMessage("APIKey not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		authorizationMode: go_google.AuthorizationModeAPIKey,
		id:                serviceConfig.APIKey,
		apiKey:            &serviceConfig.APIKey,
		httpService:       httpService,
	}, nil
}

type ServiceWithOAuth2Config struct {
	ClientID     string
	ClientSecret string
	TokenSource  tokensource.TokenSource
}

func NewServiceWithOAuth2(serviceConfig *ServiceWithOAuth2Config, bigQueryService *bigquery.Service) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ClientID == "" {
		return nil, errortools.ErrorMessage("ClientID not provided")
	}

	if serviceConfig.ClientSecret == "" {
		return nil, errortools.ErrorMessage("ClientSecret not provided")
	}

	googleServiceConfig := go_google.ServiceConfig{
		APIName:      apiName,
		ClientID:     serviceConfig.ClientID,
		ClientSecret: serviceConfig.ClientSecret,
		TokenSource:  serviceConfig.TokenSource,
	}

	googleService, e := go_google.NewService(&googleServiceConfig, bigQueryService)
	if e != nil {
		return nil, e
	}

	return &Service{
		authorizationMode: go_google.AuthorizationModeOAuth2,
		id:                go_google.ClientIDShort(serviceConfig.ClientID),
		googleService:     googleService,
	}, nil
}

type ServiceWithAccessTokenConfig struct {
	ClientID    string
	AccessToken string
}

func NewServiceWithAccessToken(serviceConfig *ServiceWithAccessTokenConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.AccessToken == "" {
		return nil, errortools.ErrorMessage("AccessToken not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		authorizationMode: go_google.AuthorizationModeAccessToken,
		accessToken:       &serviceConfig.AccessToken,
		id:                go_google.ClientIDShort(serviceConfig.ClientID),
		httpService:       httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	var request *http.Request
	var response *http.Response
	var e *errortools.Error

	if service.authorizationMode == go_google.AuthorizationModeOAuth2 {
		request, response, e = service.googleService.HTTPRequest(requestConfig)
	} else {
		// add error model
		errorResponse := go_google.ErrorResponse{}
		requestConfig.ErrorModel = &errorResponse

		if service.authorizationMode == go_google.AuthorizationModeAPIKey {
			// add api key
			requestConfig.SetParameter("key", *service.apiKey)
		}
		if service.accessToken != nil {
			// add accesstoken to header
			header := http.Header{}
			header.Set("Authorization", fmt.Sprintf("Bearer %s", *service.accessToken))
			requestConfig.NonDefaultHeaders = &header
		}

		request, response, e = service.httpService.HTTPRequest(requestConfig)

		if e != nil {
			if errorResponse.Error.Message != "" {
				e.SetMessage(errorResponse.Error.Message)
			}
		}
	}

	if e != nil {
		return request, response, e
	}

	return request, response, nil
}

func (service *Service) urlAnalytics(path string) string {
	return fmt.Sprintf("%s/%s", apiURLAnalytics, path)
}

func (service *Service) urlReporting(path string) string {
	return fmt.Sprintf("%s/%s", apiURLReporting, path)
}

func (service *Service) InitToken(scope string, accessType *string, prompt *string, state *string) *errortools.Error {
	return service.googleService.InitToken(scope, accessType, prompt, state)
}

func (service *Service) APIName() string {
	return apiName
}

func (service *Service) APIKey() string {
	return service.id
}

func (service *Service) APICallCount() int64 {
	if service.googleService == nil {
		return 0
	}
	return service.googleService.APICallCount()
}

func (service *Service) APIReset() {
	if service.googleService == nil {
		return
	}
	service.googleService.APIReset()
}
