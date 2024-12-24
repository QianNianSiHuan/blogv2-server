package conf

import "fmt"

type ES struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	IsHttps  bool   `yaml:"is_Https"`
	Password string `yaml:"password"`
}

func (e ES) Url() string {
	if e.IsHttps {
		return fmt.Sprintf("https://%s", e.Addr)
	} else {
		return fmt.Sprintf("http://%s", e.Addr)
	}
}
