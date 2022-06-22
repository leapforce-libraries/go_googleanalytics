package googleanalytics

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	g_types "github.com/leapforce-libraries/go_googleanalytics/types"
	go_http "github.com/leapforce-libraries/go_http"
)

type Account4 struct {
	Name        string                 `json:"name"`
	CreateTime  g_types.DateTimeString `json:"createTime"`
	UpdateTime  g_types.DateTimeString `json:"updateTime"`
	DisplayName string                 `json:"displayName"`
	RegionCode  string                 `json:"regionCode"`
	Deleted     bool                   `json:"deleted"`
}

type GetAccountConfig struct {
	AccountId string
}

func (service *Service) GetAccount(config *GetAccountConfig) (*Account4, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetAccountConfig must not be nil")
	}

	account := Account4{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.urlAdmin(fmt.Sprintf("accounts/%s", config.AccountId)),
		ResponseModel: &account,
	}

	_, _, e := service.googleService().HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &account, nil
}
