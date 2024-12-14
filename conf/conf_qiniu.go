package conf

type QiNiu struct {
	Enable    bool   `yaml:"enable" json:"enable"`
	AccessKey string `yaml:"accessKey" json:"accessKey"`
	SecretKey string `yaml:"secretKey" json:"secretKey"`
	Bucket    string `yaml:"bucket" json:"bucket"`
	Uri       string `yaml:"uri" json:"uri"`
	Region    string `yaml:"region" json:"region"`
	Prefix    string `yaml:"prefix" json:"prefix"`
	Size      int    `yaml:"size" json:"size"`
	Expiry    int    `yaml:"expiry" json:"expiry"` //过期时间
}
