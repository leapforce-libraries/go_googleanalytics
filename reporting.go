package googleanalyticsreporting

import (
	"context"
	"errors"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	"golang.org/x/oauth2"
	"google.golang.org/api/analyticsreporting/v4"
	"google.golang.org/api/option"
)

const (
	APIName string = "GoogleAnalyticsReporting"
)

// GoogleAnalyticsReporting stores GoogleAnalyticsReporting configuration
//
type ReportingService struct {
	googleClient     *google.GoogleClient
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
func NewReportingService(clientID string, clientSecret string, scope string, bigQuery *google.BigQuery) (*ReportingService, *errortools.Error) {
	config := google.GoogleClientConfig{
		APIName:      APIName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	googleClient := google.NewGoogleClient(config, bigQuery)

	tokenSource := TokenSource{googleClient}

	reportingService, err := analyticsreporting.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	return &ReportingService{googleClient, reportingService}, nil
}

func (rs *ReportingService) InitToken() *errortools.Error {
	return rs.googleClient.InitToken()
}
