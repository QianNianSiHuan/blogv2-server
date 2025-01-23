package models

type BannerModel struct {
	Model
	Show  bool   `json:"show"`
	Cover string `json:"cover"`
	Href  string `json:"href"`
	Type  int8   `json:"type"` //1banner2独家推广
}
