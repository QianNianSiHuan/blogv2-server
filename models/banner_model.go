package models

type BannerModel struct {
	Model
	Cover string `json:"cover"`
	Href  string `json:"href"`
}
