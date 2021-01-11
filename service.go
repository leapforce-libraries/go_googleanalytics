package googleanalytics

import (
	"context"
	"encoding/json"
	"errors"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	"golang.org/x/oauth2"
	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/analyticsreporting/v4"
	"google.golang.org/api/option"
)

const (
	APIName string = "GoogleAnalytics"
)

// Service
//
type Service struct {
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

// methods
//
func NewService(clientID string, clientSecret string, scope string, bigQuery *google.BigQuery) (*Service, *errortools.Error) {
	config := google.ServiceConfig{
		APIName:      APIName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	googleService := google.NewService(config, bigQuery)

	tokenSource := TokenSource{googleService}

	analyticsService, err := analytics.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	reportingService, err := analyticsreporting.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	return &Service{googleService, analyticsService, reportingService}, nil
}

func NewServiceJSON(credentials *google.CredentialsJSON, bigQuery *google.BigQuery) (*Service, *errortools.Error) {
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

func (service *Service) InitToken() *errortools.Error {
	if service.googleService == nil {
		return errortools.ErrorMessage("GoogleService not initialized.")
	}

	return service.googleService.InitToken()
}
