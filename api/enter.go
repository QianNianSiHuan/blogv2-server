package api

import (
	"blogv2/api/banner_api"
	"blogv2/api/image_api"
	"blogv2/api/log_api"
	"blogv2/api/site_api"
)

type Api struct {
	SiteAPi   site_api.SiteApi
	LogApi    log_api.LogApi
	ImageApi  image_api.ImageApi
	BannerApi banner_api.BannerApi
}

var App = Api{}
