package googleanalytics

import (
	"errors"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
	"golang.org/x/oauth2"
)

const (
	apiName         string = "GoogleAnalytics"
	apiURLReporting string = "https://analyticsreporting.googleapis.com/v4"
	apiURLAnalytics string = "https://www.googleapis.com/analytics/v3"
)

// Service
//
type Service struct {
	clientID      string
	googleService *google.Service
}

type TokenSource struct {
	googleService *google.Service
}

func (tokenSource TokenSource) Token() (*oauth2.Token, error) {
	t, e := tokenSource.googleService.ValidateToken()
	if e != nil {
		return nil, errors.New(e.Message())
	}

	return &oauth2.Token{
		AccessToken:  *t.AccessToken,
		TokenType:    *t.TokenType,
		RefreshToken: *t.RefreshToken,
		Expiry:       *t.Expiry,
	}, nil
}

type ServiceConfig struct {
	ClientID     string
	ClientSecret string
}

func NewService(serviceConfig *ServiceConfig, bigQueryService *bigquery.Service) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ClientID == "" {
		return nil, errortools.ErrorMessage("ClientID not provided")
	}

	if serviceConfig.ClientSecret == "" {
		return nil, errortools.ErrorMessage("ClientSecret not provided")
	}

	googleServiceConfig := google.ServiceConfig{
		APIName:      apiName,
		ClientID:     serviceConfig.ClientID,
		ClientSecret: serviceConfig.ClientSecret,
	}

	googleService, e := google.NewService(&googleServiceConfig, bigQueryService)
	if e != nil {
		return nil, e
	}

	return &Service{
		clientID:      serviceConfig.ClientID,
		googleService: googleService,
	}, nil
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
	return service.clientID
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
