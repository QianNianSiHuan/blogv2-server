package conf

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	Debug    bool   `yaml:"debug"`  //打印全部日志
	Source   string `yaml:"source"` //源
}

func (d DB) DSN() string {
	return spew.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.DB)
}
func (d DB) Empty() bool {
	return d.User == "" && d.Password == "" && d.Host == "" && d.Port == 0
}
func (d DB) Addr() string {
	return fmt.Sprintf("%s:%d", d.Host, d.Port)
}
