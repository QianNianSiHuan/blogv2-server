package api

import "blogv2/api/site_api"

type Api struct {
	SiteAPi site_api.SiteApi
}

var App = Api{}
