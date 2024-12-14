package conf

type Upload struct {
	Size      int64    `yaml:"size"`
	WhiteList []string `yaml:"whiteList"`
	UploadDir string   `yaml:"uploadDir"`
}
