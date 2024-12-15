package models

type BannerModel struct {
	Model
	Show  bool   `json:"show"`
	Cover string `json:"cover"`
	Href  string `json:"href"`
}
