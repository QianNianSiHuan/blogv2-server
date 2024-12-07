package conf

type Jwt struct {
	Expire int    `yaml:"expire"`
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
}
