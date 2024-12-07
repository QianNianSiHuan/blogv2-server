package conf

type Email struct {
	Domain       string `yaml:"domain" json:"domain"`
	Port         int    `yaml:"port" json:"port"`
	SendEmail    string `yaml:"sendEmail" json:"sendEmail"`
	AuthCode     string `yaml:"authCode" json:"authCode"`
	SendNickName string `yaml:"sendNickName" json:"sendNickName"`
	SSL          bool   `yaml:"ssl" json:"ssl"`
	TLS          bool   `yaml:"tls" json:"tls"`
}
