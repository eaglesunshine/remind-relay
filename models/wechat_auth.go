package models

import (
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

const (
	OAuthUserGenderFemale int = iota
	OAuthUserGenderMale
	OAuthUserGenderUnknown
	OAuthUserPlatformWechat = "wechat"
)

// OAuthInterface 微信授权interface
type OAuthInterface interface {
	Authorize(state, callback string) string
	AccessToken(code, state, callback string) (OAuthAccessToken, error)
	Userinfo(accessToken, openID string) (OAuthUserinfo, error)
}

// OAuthAccessToken 微信授权登录token
type OAuthAccessToken struct {
	AccessToken string `json:"access_token"`
	OpenID      string `json:"openid"`
	UnionID     string `json:"unionid"`
}

// OAuthUserinfo 微信登陆用户信息
type OAuthUserinfo struct {
	Platform string `json:"platform"`
	Name     string `json:"name"`
	OpenID   string `json:"openid"`
	UnionID  string `json:"unionid"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	Detail   string `json:"-"`
}

// OAuthWechat 微信授权信息
type OAuthWechat struct {
	ClientID     string
	ClientSecret string
}

// NewOAuthWechat 创建微信授权类
func NewOAuthWechat(id, secret string) *OAuthWechat {
	return &OAuthWechat{
		ClientID:     id,
		ClientSecret: secret,
	}
}

// Authorize 授权跳转
func (o *OAuthWechat) Authorize(state, callback string) string {
	return fmt.Sprintf(
		`https://open.weixin.qq.com/connect/qrconnect?appid=%s`+
			`&response_type=code&scope=snsapi_login&state=%s&redirect_uri=%s`,
		o.ClientID,
		state,
		callback,
	)
}

// AccessToken 获取登录token
func (o *OAuthWechat) AccessToken(code, state, callback string) (OAuthAccessToken, error) {
	rtn := OAuthAccessToken{}
	resp, err := grequests.Post(
		`https://api.weixin.qq.com/sns/oauth2/access_token`,
		&grequests.RequestOptions{
			Params: map[string]string{
				"appid":      o.ClientID,
				"secret":     o.ClientSecret,
				"code":       code,
				"grant_type": "authorization_code",
			},
		},
	)
	if err != nil {
		log.Println(err)
		return rtn, err
	}
	log.Println(resp.String())

	respData := struct {
		ErrorCode int64  `json:"errcode,omitempty"`
		ErrMsg    string `json:"errmsg,omitempty"`
		// success
		AccessToken  string `json:"access_token,omitempty"`
		ExpiresIn    int64  `json:"expires_in,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
		OpenID       string `json:"openid,omitempty"`
		Scope        string `json:"scope,omitempty"`
		UnionID      string `json:"unionid,omitempty"`
	}{}

	if err := resp.JSON(&respData); err != nil {
		log.Println(err)
		return rtn, err
	}

	if respData.ErrorCode > 0 {
		return rtn, errors.New(respData.ErrMsg)
	}

	rtn.AccessToken = respData.AccessToken
	rtn.OpenID = respData.OpenID
	rtn.UnionID = respData.UnionID
	return rtn, nil
}

// Userinfo 获取用户信息
func (o *OAuthWechat) Userinfo(accessToken, openID string) (OAuthUserinfo, error) {
	rtn := OAuthUserinfo{
		Platform: OAuthUserPlatformWechat,
		Gender:   OAuthUserGenderUnknown,
	}
	resp, err := grequests.Get(
		`https://api.weixin.qq.com/sns/userinfo`,
		&grequests.RequestOptions{
			Params: map[string]string{
				"access_token": accessToken,
				"openid":       openID,
				"lang":         "zh_CN",
			},
		},
	)
	if err != nil {
		log.Println(err)
		return rtn, err
	}

	rtn.Detail = resp.String()
	log.Println(rtn.Detail)

	respData := struct {
		ErrorCode  int64    `json:"errorcode,omitempty"`
		ErrorMsg   string   `json:"errormsg,omitempty"`
		OpenID     string   `json:"openid,omitempty"`
		Nickname   string   `json:"nickname,omitempty"`
		Sex        int      `json:"sex,omitempty"`
		Province   string   `json:"province,omitempty"`
		City       string   `json:"city,omitempty"`
		Country    string   `json:"country,omitempty"`
		HeadImgURL string   `json:"headimgurl,omitempty"`
		Privilege  []string `json:"privilege,omitempty"`
		UnionID    string   `json:"unionid,omitempty"`
	}{}

	if err := resp.JSON(&respData); err != nil {
		log.Println(err)
		return rtn, err
	}

	if respData.ErrorCode != 0 {
		log.Println(respData.ErrorCode)
		return rtn, errors.New(respData.ErrorMsg)
	}

	rtn.OpenID = respData.OpenID
	rtn.UnionID = respData.UnionID
	rtn.Name = respData.Nickname
	rtn.Avatar = respData.HeadImgURL

	switch respData.Sex {
	case 2:
		rtn.Gender = OAuthUserGenderFemale
		break
	case 1:
		rtn.Gender = OAuthUserGenderMale
		break
	default:
		rtn.Gender = OAuthUserGenderUnknown
		break
	}

	return rtn, err
}
