package conf

import "blogv2/conf/site"

type Site struct {
	SiteInfo   site.SiteInfo   `yaml:"siteInfo" json:"siteInfo"`
	Project    site.Project    `yaml:"project" json:"project"`
	Seo        site.Seo        `yaml:"seo" json:"seo"`
	About      site.About      `yaml:"about" json:"about"`
	Login      site.Login      `yaml:"login" json:"login"`
	IndexRight site.IndexRight `yaml:"indexRight" json:"indexRight"`
	Article    site.Article    `yaml:"article" json:"article"`
}
