package googleanalytics

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
	credentials "github.com/leapforce-libraries/go_google/credentials"
	"golang.org/x/oauth2"
	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/analyticsreporting/v4"
	"google.golang.org/api/option"
)

const (
	apiName string = "GoogleAnalytics"
	apiURL  string = "https://analyticsreporting.googleapis.com/v4"
)

// Service
//
type Service struct {
	clientID         string
	googleService    *google.Service
	AnalyticsService *analytics.Service
	ReportingService *analyticsreporting.Service
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

	tokenSource := TokenSource{googleService}

	analyticsService, err := analytics.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	reportingService, err := analyticsreporting.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	return &Service{
		clientID:         serviceConfig.ClientID,
		googleService:    googleService,
		AnalyticsService: analyticsService,
		ReportingService: reportingService,
	}, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURL, path)
}

func NewServiceJSON(credentials *credentials.CredentialsJSON) (*Service, *errortools.Error) {
	if credentials == nil {
		return nil, errortools.ErrorMessage("Credentials can be not be a nil pointer.")
	}

	credentialsJSON, err := json.Marshal(&credentials)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	analyticsService, err := analytics.NewService(context.Background(), option.WithCredentialsJSON(credentialsJSON))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	reportingService, err := analyticsreporting.NewService(context.Background(), option.WithCredentialsJSON(credentialsJSON))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	return &Service{
		AnalyticsService: analyticsService,
		ReportingService: reportingService,
	}, nil
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
