package googleanalytics

import (
	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type MetricType string

const (
	MetricTypeUnspecified MetricType = "METRIC_TYPE_UNSPECIFIED"
	MetricTypeInteger     MetricType = "INTEGER"
	MetricTypeFloat       MetricType = "FLOAT"
	MetricTypeCurrency    MetricType = "CURRENCY"
	MetricTypePercent     MetricType = "PERCENT"
	MetricTypeTime        MetricType = "TIME"
)

type ReportRequestBody struct {
	ReportRequests    []*ReportRequest `json:"reportRequests"`
	UseResourceQuotas *bool            `json:"useResourceQuotas,omitempty"`
}

type ReportRequest struct {
	ViewID     string       `json:"viewId"`
	DateRanges *[]DateRange `json:"dateRanges,omitempty"`
	//samplingLevel
	Dimensions []Dimension `json:"dimensions,omitempty"`
	Metrics    []Metric    `json:"metrics,omitempty"`
	//metricFilterClauses
	//filtersExpression
	//orderBys
	//segments
	//pivots
	//cohortGroup
	PageToken        *string `json:"pageToken,omitempty"`
	PageSize         *int64  `json:"pageSize,omitempty"`
	IncludeEmptyRows *bool   `json:"includeEmptyRows,omitempty"`
	HideTotals       *bool   `json:"hideTotals,omitempty"`
	HideValueRanges  *bool   `json:"hideValueRanges,omitempty"`
}

type DateRange struct {
	StartDate civil.Date `json:"startDate"`
	EndDate   civil.Date `json:"endDate"`
}

type Dimension struct {
	Name string `json:"name"`
	//histogramBuckets
}

type Metric struct {
	Expression     string      `json:"expression"`
	Alias          string      `json:"alias,omitempty"`
	FormattingType *MetricType `json:"formattingType,omitempty"`
}

type ReportResponseBody struct {
	Reports   []ReportResponse `json:"reports"`
	QueryCost int64            `json:"queryCost"`
	//resourceQuotasRemaining
}

type ReportResponse struct {
	ColumnHeader  ColumnHeader `json:"columnHeader"`
	Data          ReportData   `json:"data"`
	NextPageToken string       `json:"nextPageToken"`
}

type ColumnHeader struct {
	Dimensions   []string     `json:"dimensions"`
	MetricHeader MetricHeader `json:"metricHeader"`
}

type MetricHeader struct {
	MetricHeaderEntries []MetricHeaderEntry `json:"metricHeaderEntries"`
	//pivotHeaders
}

type MetricHeaderEntry struct {
	Name string     `json:"name"`
	Type MetricType `json:"type"`
}

type ReportData struct {
	Rows     []ReportRow       `json:"rows"`
	Totals   []DateRangeValues `json:"totals"`
	RowCount int64             `json:"rowCount"`
}

type ReportRow struct {
	Dimensions []string          `json:"dimensions"`
	Metrics    []DateRangeValues `json:"metrics"`
}

type DateRangeValues struct {
	Values []string `json:"values"`
	//pivotValueRegions
}

func (service *Service) BatchGet(reportRequestBody *ReportRequestBody) (*ReportResponseBody, *errortools.Error) {
	reportResponseBody := ReportResponseBody{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url("reports:batchGet"),
		BodyModel:     reportRequestBody,
		ResponseModel: &reportResponseBody,
	}

	_, _, e := service.googleService.Post(&requestConfig)
	return &reportResponseBody, e
}

/*
func (service *Service) BatchGet(reportRequestBody *ReportRequestBody) (*ColumnHeader, *ReportData, *errortools.Error) {
	reportResponse, e := service.batchGet(reportRequestBody)
	if e != nil {
		return nil, nil, e
	}

	return &reportResponse.ColumnHeader, &reportResponse.Data, nil
}

func (service *Service) BatchGetRowCount(reportRequestBody *ReportRequestBody) (int64, *errortools.Error) {
	reportResponse, e := service.batchGet(reportRequestBody)
	if e != nil {
		return 0, e
	}

	return reportResponse.Data.RowCount, nil
}*/
