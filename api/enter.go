package api

import (
	"blogv2/api/ai_api"
	"blogv2/api/article_api"
	"blogv2/api/banner_api"
	"blogv2/api/captcha_api"
	"blogv2/api/comment_api"
	"blogv2/api/data_api"
	"blogv2/api/feed_back_api"
	"blogv2/api/image_api"
	"blogv2/api/log_api"
	"blogv2/api/search_api"
	"blogv2/api/site_api"
	"blogv2/api/user_api"
)

type Api struct {
	SiteAPi     site_api.SiteApi
	LogApi      log_api.LogApi
	ImageApi    image_api.ImageApi
	BannerApi   banner_api.BannerApi
	CaptchaApi  captcha_api.CaptchaApi
	UserApi     user_api.UserApi
	ArticleApi  article_api.ArticleApi
	CommentApi  comment_api.CommentApi
	SearchApi   search_api.SearchApi
	AiApi       ai_api.AiApi
	FeedBackApi feed_back_api.FeedBackApi
	DataApi     data_api.DataApi
}

var App = Api{}
