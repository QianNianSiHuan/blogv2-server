package conf

import "fmt"

type QQ struct {
	AppID    string `yaml:"appID" json:"appID"`
	AppKey   string `yaml:"appKey" json:"appKey"`
	Redirect string `yaml:"redirect" json:"redirect"`
}

func (q QQ) Url() string {
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
