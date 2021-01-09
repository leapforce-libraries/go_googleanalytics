package googleanalytics

import (
	"context"
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
	googleClient     *google.GoogleClient
	AnalyticsService *analytics.Service
	ReportingService *analyticsreporting.Service
}

type TokenSource struct {
	googleClient *google.GoogleClient
}

func (tokenSource TokenSource) Token() (*oauth2.Token, error) {
	t, e := tokenSource.googleClient.ValidateToken()
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
func NewService(clientID string, clientSecret string, bigQuery *google.BigQuery) (*Service, *errortools.Error) {
	config := google.GoogleClientConfig{
		APIName:      APIName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	googleClient := google.NewGoogleClient(config, bigQuery)

	tokenSource := TokenSource{googleClient}

	analyticsService, err := analytics.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	reportingService, err := analyticsreporting.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	return &Service{googleClient, analyticsService, reportingService}, nil
}

func (rs *Service) InitToken() *errortools.Error {
	return rs.googleClient.InitToken()
}
