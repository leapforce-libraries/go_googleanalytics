package googleanalytics

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	data "github.com/leapforce-libraries/go_googleanalytics/data"
	go_http "github.com/leapforce-libraries/go_http"
)

func (service *Service) RunReport(propertyId string, reportRequest *data.ReportRequest) (*data.RunReportResponse, *errortools.Error) {
	runReportResponse := data.RunReportResponse{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.urlData(fmt.Sprintf("properties/%s:runReport", propertyId)),
		BodyModel:     reportRequest,
		ResponseModel: &runReportResponse,
	}

	_, _, e := service.googleService().HttpRequest(&requestConfig)
	return &runReportResponse, e
}
