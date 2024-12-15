package site

type SitesInfo struct {
	Title string `yaml:"title" json:"title"`
	Logo  string `yaml:"logo" json:"logo"`
	Beian string `yaml:"beian" json:"beian"`
	Mode  int8   `yaml:"mode" json:"mode"` //1.社区模式2.博客模式
}
type Project struct {
	Title   string `yaml:"title" json:"title"`
	Icon    string `yaml:"icon" json:"icon"`
	WebPath string `yaml:"webPath" json:"webPath"`
}
type Seo struct {
	Keywords    string `yaml:"keywords" json:"keywords"`
	Description string `yaml:"description" json:"description"`
}
type About struct {
	SiteDate string `yaml:"siteDate" json:"siteDate"`
	QQ       string `yaml:"QQ" json:"QQ"`
	Wechat   string `yaml:"wechat" json:"wechat"`
	Version  string `yaml:"-" json:"version"`
	Gitee    string `yaml:"gitee" json:"gitee"`
	Bilibili string `yaml:"bilibili" json:"bilibili"`
	GitHub   string `yaml:"gitHub" json:"gitHub"`
}

type Login struct {
	QQLogin          bool `yaml:"qqLogin" json:"qqLogin"`
	UsernamePwdLogin bool `yaml:"usernamePwdLogin" json:"usernamePwdLogin"`
	EmailLogin       bool `yaml:"emailLogin" json:"emailLogin"`
	Captcha          bool `yaml:"captcha" json:"captcha"`
}
type ComponentInfo struct {
	Title  string `yaml:"title" json:"title"`
	Enable bool   `yaml:"enable" json:"enable"`
}
type IndexRight struct {
	List []ComponentInfo
}
type Article struct {
	NoExamine bool `yaml:"noExamine" json:"noExamine"` //免审核
}
