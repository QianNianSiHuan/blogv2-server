package site_api

import (
	"blogv2/common/res"
	"blogv2/conf"
	"blogv2/core"
	"blogv2/global"
	"blogv2/middleware"
	"blogv2/service/redis_service/redis_site"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

type SiteApi struct {
}

type SiteInfoRequest struct {
	Name string `uri:"name"`
}

type QiNiu struct {
	Enable bool `json:"enable"`
}
type Ai struct {
	Enable bool `json:"enable"`
}

type SiteInfoResponse struct {
	conf.Site
	QiNiu QiNiu `json:"qiNiu"`
	Ai    Ai    `json:"ai"`
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(c, err)
	}
	if cr.Name == "site" {
		redis_site.SetFlow()
		global.Config.Site.About.Version = global.Version
		res.SuccessWithData(c, SiteInfoResponse{
			Site: global.Config.Site,
			QiNiu: QiNiu{
				Enable: global.Config.QiNiu.Enable,
			},
			Ai: Ai{
				Enable: global.Config.Ai.Enable,
			},
		})
		return
	}
	//判断角色是否管理员
	middleware.AdminMiddleware(c)
	_, ok := c.Get("claims")
	if !ok {
		return
	}
	var data any
	switch cr.Name {
	case "email":
		rep := global.Config.Email
		rep.AuthCode = "******"
		data = rep
	case "qq":
		rep := global.Config.QQ
		rep.AppKey = "******"
		data = rep

	case "ai":
		rep := global.Config.Ai
		rep.SecretKey = "******"
		data = rep
	case "qiNiu":
		rep := global.Config.QiNiu
		rep.SecretKey = "******"
		data = rep
	default:
		res.FailWithMsg(c, "不存在的配置")
		return
	}
	res.SuccessWithData(c, data)
}
func (SiteApi) SiteInfoQQView(c *gin.Context) {
	res.SuccessWithData(c, global.Config.QQ.Url())
}

type SiteUpdateRequest struct {
	Name string `uri:"name" binding:"required" label:"姓名"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	//log := log_service.GetLog(c)
	var cr SiteUpdateRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	var rep any
	switch cr.Name {
	case "site":
		var data conf.Site
		err = c.ShouldBindJSON(&data)
		rep = data
	case "email":
		var data conf.Email
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qq":
		var data conf.QQ
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qiNiu":
		var data conf.QiNiu
		err = c.ShouldBindJSON(&data)
		rep = data
	case "ai":
		var data conf.Ai
		err = c.ShouldBindJSON(&data)
		rep = data
	default:
		res.FailWithMsg(c, "不存在的配置")
		return
	}
	if err != nil {
		res.FailWithError(c, err)
		return
	}
	switch s := rep.(type) {
	case conf.Site:
		//判断站点信息更新前端部分
		global.Config.Site = s
		err = UpdateSite(s)
		if err != nil {
			res.FailWithError(c, err)
			return
		}
	case conf.Email:
		if s.AuthCode == "******" {
			s.AuthCode = global.Config.Email.AuthCode
		}
		global.Config.Email = s
	case conf.QQ:
		if s.AppKey == "******" {
			s.AppKey = global.Config.QQ.AppKey
		}
		global.Config.QQ = s
	case conf.QiNiu:
		if s.SecretKey == "******" {
			s.SecretKey = global.Config.QiNiu.SecretKey
		}
		global.Config.QiNiu = s
	case conf.Ai:
		if s.SecretKey == "******" {
			s.SecretKey = global.Config.Ai.SecretKey
		}
		global.Config.Ai = s
	}
	//更改配置文件
	core.SetConf()
	res.SuccessWithMsg(c, "更新站点配置成功")
	return
}

func UpdateSite(site conf.Site) (err error) {
	if site.Project.Icon == "" && site.Project.Title == "" && site.Project.WebPath == "" &&
		site.Seo.Keywords == "" && site.Seo.Description == "" {
		return nil
	}
	if site.Project.WebPath == "" {
		return errors.New("请配置前端地址")
	}
	file, err := os.Open(site.Project.WebPath)
	if err != nil {
		return errors.New(fmt.Sprintf("%s 文件不存在", site.Project.WebPath))
	}
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		logrus.Errorf("goquery解析失败 %s", err)
		return errors.New("文件解析失败")
	}
	if site.Project.Title != "" {
		doc.Find("title").SetText(site.Project.Title)
	}
	if site.Project.Icon != "" {
		selection := doc.Find("link[rel=\"icon\"]")
		if selection.Length() > 0 {
			selection.SetAttr("href", site.Project.Icon)
		} else {
			// 没有就创建
			doc.Find("head").AppendHtml(fmt.Sprintf("<link rel=\"icon\" href=\"%s\">", site.Project.Icon))
		}
	}
	if site.Seo.Keywords != "" {
		selection := doc.Find("meta[name=\"keywords\"]")
		if selection.Length() > 0 {
			selection.SetAttr("content", site.Seo.Keywords)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"keywords\" content=\"%s\">", site.Seo.Keywords))
		}
	}
	if site.Seo.Description != "" {
		selection := doc.Find("meta[name=\"description\"]")
		if selection.Length() > 0 {
			selection.SetAttr("content", site.Seo.Description)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", site.Seo.Description))
		}
	}

	html, err := doc.Html()
	if err != nil {
		logrus.Errorf("生成html失败 %s", err)
		return errors.New("生成html失败")
	}

	err = os.WriteFile(site.Project.WebPath, []byte(html), 0666)
	if err != nil {
		logrus.Errorf("文件写入失败 %s", err)
		return errors.New("文件写入失败")
	}
	return nil
}
