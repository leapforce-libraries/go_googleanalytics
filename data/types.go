package data

type ReportRequest struct {
	Dimensions          []Dimension          `json:"dimensions,omitempty"`
	Metrics             []Metric             `json:"metrics,omitempty"`
	DateRanges          *[]DateRange         `json:"dateRanges,omitempty"`
	DimensionFilter     *FilterExpression    `json:"dimensionFilter,omitempty"`
	MetricFilter        *FilterExpression    `json:"metricFilter,omitempty"`
	Offset              *string              `json:"offset,omitempty"`
	Limit               *string              `json:"limit,omitempty"`
	MetricAggregations  *[]MetricAggregation `json:"metricAggregations,omitempty"`
	OrderBys            *[]OrderBy           `json:"orderBys,omitempty"`
	CurrencyCode        *string              `json:"currencyCode,omitempty"`
	CohortSpec          *CohortSpec          `json:"cohortSpec,omitempty"`
	KeepEmptyRows       *bool                `json:"keepEmptyRows,omitempty"`
	ReturnPropertyQuota *bool                `json:"returnPropertyQuota,omitempty"`
}

type Dimension struct {
	Name                string               `json:"name"`
	DimensionExpression *DimensionExpression `json:"dimensionExpression,omitempty"`
}

type DimensionExpression struct {
	LowerCase   *CaseExpression        `json:"lowerCase,omitempty"`
	UpperCase   *CaseExpression        `json:"upperCase,omitempty"`
	Concatenate *ConcatenateExpression `json:"concatenate,omitempty"`
}

type CaseExpression struct {
	DimensionName string `json:"dimensionName"`
}

type ConcatenateExpression struct {
	DimensionNames []string `json:"dimensionNames"`
	Delimiter      string   `json:"delimiter"`
}

type Metric struct {
	Name       string  `json:"name"`
	Expression *string `json:"expression,omitempty"`
	Invisible  *bool   `json:"invisible,omitempty"`
}

type DateRange struct {
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Name      *string `json:"name,omitempty"`
}

type FilterExpression struct {
	AndGroup      *FilterExpressionList `json:"andGroup,omitempty"`
	OrGroup       *FilterExpressionList `json:"orGroup,omitempty"`
	NotExpression *FilterExpression     `json:"notExpression,omitempty"`
	Filter        *Filter               `json:"filter,omitempty"`
}

type FilterExpressionList struct {
	Expressions []FilterExpression `json:"expressions"`
}

type Filter struct {
	FieldName     string         `json:"fieldName"`
	StringFilter  *StringFilter  `json:"stringFilter,omitempty"`
	InListFilter  *InListFilter  `json:"inListFilter,omitempty"`
	NumericFilter *NumericFilter `json:"numericFilter,omitempty"`
	BetweenFilter *BetweenFilter `json:"betweenFilter,omitempty"`
}

type StringFilter struct {
	MatchType     string `json:"matchType"`
	Value         string `json:"value"`
	CaseSensitive bool   `json:"caseSensitive"`
}

type InListFilter struct {
	Values        []string `json:"values"`
	CaseSensitive bool     `json:"caseSensitive"`
}

type NumericFilter struct {
	Operation Operation    `json:"operation"`
	Value     NumericValue `json:"values"`
}

type Operation string

const (
	OperationUnspecified        Operation = "OPERATION_UNSPECIFIED"
	OperationEqual              Operation = "EQUAL"
	OperationLessThan           Operation = "LESS_THAN"
	OperationLessThanOrEqual    Operation = "LESS_THAN_OR_EQUAL"
	OperationGreaterThan        Operation = "GREATER_THAN"
	OperationGreaterThanOrEqual Operation = "GREATER_THAN_OR_EQUAL"
)

type NumericValue struct {
	Int64Value  int64   `json:"int64Value,omitempty"`
	DoubleValue float64 `json:"doubleValue,omitempty"`
}

type BetweenFilter struct {
	FromValue NumericValue `json:"fromValue"`
	ToValue   NumericValue `json:"toValue"`
}

type MetricAggregation string

const (
	MetricAggregationUnspecified Operation = "METRIC_AGGREGATION_UNSPECIFIED"
	MetricAggregationTotal       Operation = "TOTAL"
	MetricAggregationMinimum     Operation = "MINIMUM"
	MetricAggregationMaximum     Operation = "MAXIMUM"
	MetricAggregationCount       Operation = "COUNT"
)

type OrderBy struct {
	Desc      *bool          `json:"desc,omitempty"`
	Metric    *MetricOrderBy `json:"metric,omitempty"`
	Dimension *NumericValue  `json:"dimension,omitempty"`
	Pivot     *NumericValue  `json:"pivot,omitempty"`
}

type MetricOrderBy struct {
	MetricName string `json:"metricName"`
}

type DimensionOrderBy struct {
	DimensionName string    `json:"dimensionName"`
	OrderType     OrderType `json:"orderType"`
}

type OrderType string

const (
	OrderTypeUnspecified                 Operation = "ORDER_TYPE_UNSPECIFIED"
	OrderTypeAlphanumeric                Operation = "ALPHANUMERIC"
	OrderTypeCaseInsensitiveAlphanumeric Operation = "CASE_INSENSITIVE_ALPHANUMERIC"
	OrderTypeNumeric                     Operation = "NUMERIC"
)

type PivotOrderBy struct {
	MetricName      string           `json:"metricName"`
	PivotSelections []PivotSelection `json:"pivotSelections"`
}

type PivotSelection struct {
	DimensionName  string `json:"dimensionName"`
	DimensionValue string `json:"dimensionValue"`
}

type CohortSpec struct {
	Cohorts              *[]Cohort             `json:"cohorts"`
	CohortsRange         *[]CohortsRange       `json:"cohortsRange"`
	CohortReportSettings *CohortReportSettings `json:"cohortReportSettings"`
}

type Cohort struct {
	Name      *string   `json:"name,omitempty"`
	Dimension string    `json:"dimension"`
	DateRange DateRange `json:"dateRange"`
}

type CohortsRange struct {
	Granularity Granularity `json:"granularity,omitempty"`
	StartOffset *int64      `json:"startOffset,omitempty"`
	EndOffset   int64       `json:"endOffset"`
}

type Granularity string

const (
	GranularityUnspecified Operation = "GRANULARITY_UNSPECIFIED"
	GranularityDaily       Operation = "DAILY"
	GranularityWeekly      Operation = "WEEKLY"
	GranularityMonthly     Operation = "MONTHLY"
)

type CohortReportSettings struct {
	Accumulate bool `json:"accumulate"`
}

type RunReportResponse struct {
	DimensionHeaders []DimensionHeader `json:"dimensionHeaders"`
	MetricHeaders    []MetricHeader    `json:"metricHeaders"`
	Rows             []Row             `json:"rows"`
	Totals           []Row             `json:"totals"`
	Maximums         []Row             `json:"maximums"`
	Minimums         []Row             `json:"minimums"`
	RowCount         int64             `json:"rowCount"`
	Metadata         ResponseMetaData  `json:"metadata"`
	PropertyQuota    PropertyQuota     `json:"propertyQuota"`
	Kind             string            `json:"kind"`
}

type DimensionHeader struct {
	Name string `json:"name"`
}

type Row struct {
	DimensionValues []DimensionValue `json:"dimensionValues"`
	MetricValues    []MetricValue    `json:"metricValues"`
}

type MetricHeader struct {
	Name       string     `json:"name"`
	MetricType MetricType `json:"type"`
}

type MetricType string

const (
	MetricTypeUnspecified  MetricType = "METRIC_TYPE_UNSPECIFIED"
	MetricTypeInteger      MetricType = "TYPE_INTEGER"
	MetricTypeFloat        MetricType = "TYPE_FLOAT"
	MetricTypeSeconds      MetricType = "TYPE_SECONDS"
	MetricTypeMilliseconds MetricType = "TYPE_MILLISECONDS"
	MetricTypeMinutes      MetricType = "TYPE_MINUTES"
	MetricTypeHours        MetricType = "TYPE_HOURS"
	MetricTypeStandard     MetricType = "TYPE_STANDARD"
	MetricTypeCurrency     MetricType = "TYPE_CURRENCY"
	MetricTypeFeet         MetricType = "TYPE_FEET"
	MetricTypeMiles        MetricType = "TYPE_MILES"
	MetricTypeMeters       MetricType = "TYPE_METERS"
	MetricTypeKilometers   MetricType = "TYPE_KILOMETERS"
)

type DimensionValue struct {
	Value string `json:"value"`
}

type MetricValue struct {
	Value string `json:"value"`
}

type ResponseMetaData struct {
	DataLossFromOtherRow      bool                      `json:"dataLossFromOtherRow"`
	SchemaRestrictionResponse SchemaRestrictionResponse `json:"schemaRestrictionResponse"`
	CurrencyCode              string                    `json:"currencyCode"`
	TimeZone                  string                    `json:"timeZone"`
	EmptyReason               string                    `json:"emptyReason"`
	SubjectToThresholding     bool                      `json:"subjectToThresholding"`
}

type SchemaRestrictionResponse struct {
	ActiveMetricRestrictions []ActiveMetricRestriction `json:"activeMetricRestrictions"`
}

type ActiveMetricRestriction struct {
	RestrictedMetricTypes []ActiveMetricRestriction `json:"restrictedMetricTypes"`
	MetricName            string                    `json:"metricName"`
}

type RestrictedMetricType string

const (
	RestrictedMetricTypeUnspecified RestrictedMetricType = "RESTRICTED_METRIC_TYPE_UNSPECIFIED"
	RestrictedMetricTypeCostData    RestrictedMetricType = "COST_DATA"
	RestrictedMetricTypeRevenueData RestrictedMetricType = "REVENUE_DATA"
)

type PropertyQuota struct {
	TokensPerDay                          QuotaStatus `json:"tokensPerDay"`
	TokensPerHour                         QuotaStatus `json:"tokensPerHour"`
	ConcurrentRequests                    QuotaStatus `json:"concurrentRequests"`
	ServerErrorsPerProjectPerHour         QuotaStatus `json:"serverErrorsPerProjectPerHour"`
	PotentiallyThresholdedRequestsPerHour QuotaStatus `json:"potentiallyThresholdedRequestsPerHour"`
}

type QuotaStatus struct {
	Consumed  int64 `json:"consumed"`
	Remaining int64 `json:"remaining"`
}
