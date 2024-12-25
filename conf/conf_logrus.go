package conf

type Log struct {
	Debug bool   `yaml:"debug"`
	App   string `yaml:"app"`
	Dir   string `yaml:"dir"`
}
