package qq_service

import (
	"blogv2/global"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AccessTokenType struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	Openid       string `json:"openid"`
	RefreshToken string `json:"refresh_token"`
}
type UserInfo struct {
	Ret             int    `json:"ret"`
	Msg             string `json:"msg"`
	IsLost          int    `json:"is_lost"`
	Nickname        string `json:"nickname"`
	Gender          string `json:"gender"`
	GenderType      int    `json:"gender_type"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Year            string `json:"year"`
	Figureurl       string `json:"figureurl"`
	Figureurl1      string `json:"figureurl_1"`
	Figureurl2      string `json:"figureurl_2"`
	FigureurlQq1    string `json:"figureurl_qq_1"`
	FigureurlQq2    string `json:"figureurl_qq_2"`
	FigureurlQq     string `json:"figureurl_qq"`
	IsYellowVip     string `json:"is_yellow_vip"`
	Vip             string `json:"vip"`
	YellowVipLevel  string `json:"yellow_vip_level"`
	Level           string `json:"level"`
	IsYellowYearVip string `json:"is_yellow_year_vip"`
}
type QQUserInfo struct {
	OpenID   string
	Nickname string
	Avatar   string
}

func GetUserInfo(code string) (info QQUserInfo, err error) {
	at, err := getAccessToken(code)
	if err != nil {
		return
	}
	u, err := getUserInfo(at)
	if err != nil {
		return
	}
	return QQUserInfo{
		OpenID:   at.Openid,
		Nickname: u.Nickname,
		Avatar:   u.FigureurlQq,
	}, nil
}
func getAccessToken(code string) (at AccessTokenType, err error) {
	qq := global.Config.QQ

	baseUrl, _ := url.Parse("https://graph.qq.com/oauth2.0/token")

	p := url.Values{}
	p.Add("grant_type", "authorization_code")
	p.Add("client_id", qq.AppID)
	p.Add("client_secret", qq.AppKey)
	p.Add("code", code)
	p.Add("fmt", "json")
	p.Add("redirect_uri", qq.Redirect)
	p.Add("need_openid", "1")
	baseUrl.RawQuery = p.Encode()
	response, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}

	byteData, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(byteData, &at)
	if err != nil {
		return
	}
	if at.AccessToken == "" {
		// 获取更人性化的错误
		if strings.Contains(string(byteData), "code is reused error") {
			err = errors.New("code失效")
			return
		}
		if strings.Contains(string(byteData), "client_secret is wrong") {
			err = errors.New("secret错误")
			return
		}
		fmt.Println(string(byteData))
		err = errors.New("获取accessToken失败")
		return
	}
	return
}
func getUserInfo(at AccessTokenType) (userinfo UserInfo, err error) {
	qq := global.Config.QQ

	baseUrl, _ := url.Parse("https://graph.qq.com/user/get_user_info")

	p := url.Values{}
	p.Add("access_token", at.AccessToken)
	p.Add("oauth_consumer_key", qq.AppID)
	p.Add("openid", at.Openid)
	baseUrl.RawQuery = p.Encode()
	response, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}

	byteData, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(byteData, &userinfo)
	if err != nil {
		return
	}
	if userinfo.Ret != 0 {
		err = errors.New(userinfo.Msg)
		return
	}
	return
}
