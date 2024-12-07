package api

import (
	"blogv2/api/log_api"
	"blogv2/api/site_api"
)

type Api struct {
	SiteAPi site_api.SiteApi
	LogApi  log_api.LogApi
}

var App = Api{}
