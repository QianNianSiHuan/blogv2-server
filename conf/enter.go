package conf

var confPath = "settings-dev.yaml"

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     []DB   `yaml:"db"` //数据库连接列表
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
	Site   Site   `yaml:"site"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	Ai     Ai     `yaml:"ai"`
	QiNiu  QiNiu  `yaml:"qiNiu"`
	Upload Upload `yaml:"upload"`
	ES     ES     `yaml:"es"`
	River  River  `yaml:"river"`
}
