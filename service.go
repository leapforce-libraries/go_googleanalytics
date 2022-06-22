package googleanalytics

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_google "github.com/leapforce-libraries/go_google"
)

const (
	apiName         string = "GoogleAnalytics"
	apiURLReporting string = "https://analyticsreporting.googleapis.com/v4"
	apiURLAnalytics string = "https://www.googleapis.com/analytics/v3"
	apiURLData      string = "https://analyticsdata.googleapis.com/v1beta"
	apiURLAdmin     string = "https://analyticsadmin.googleapis.com/v1alpha"
)

type Service go_google.Service

func NewServiceWithAccessToken(cfg *go_google.ServiceWithAccessTokenConfig) (*Service, *errortools.Error) {
	googleService, e := go_google.NewServiceWithAccessToken(cfg)
	if e != nil {
		return nil, e
	}
	service := Service(*googleService)
	return &service, nil
}

func NewServiceWithApiKey(cfg *go_google.ServiceWithApiKeyConfig) (*Service, *errortools.Error) {
	googleService, e := go_google.NewServiceWithApiKey(cfg)
	if e != nil {
		return nil, e
	}
	service := Service(*googleService)
	return &service, nil
}

func NewServiceWithOAuth2(cfg *go_google.ServiceWithOAuth2Config) (*Service, *errortools.Error) {
	googleService, e := go_google.NewServiceWithOAuth2(cfg)
	if e != nil {
		return nil, e
	}
	service := Service(*googleService)
	return &service, nil
}

func (service *Service) urlAnalytics(path string) string {
	return fmt.Sprintf("%s/%s", apiURLAnalytics, path)
}

func (service *Service) urlReporting(path string) string {
	return fmt.Sprintf("%s/%s", apiURLReporting, path)
}

func (service *Service) urlData(path string) string {
	return fmt.Sprintf("%s/%s", apiURLData, path)
}

func (service *Service) urlAdmin(path string) string {
	return fmt.Sprintf("%s/%s", apiURLAdmin, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.googleService().ApiKey()
}

func (service *Service) ApiCallCount() int64 {
	return service.googleService().ApiCallCount()
}

func (service *Service) ApiReset() {
	service.googleService().ApiReset()
}

func (service *Service) googleService() *go_google.Service {
	googleService := go_google.Service(*service)
	return &googleService
}
