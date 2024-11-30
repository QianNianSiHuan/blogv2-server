package conf

var confPath = "settings.yaml"

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`  //读库
	DB1    DB     `yaml:"db1"` //写库
}
