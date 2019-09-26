/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:22:56
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-15 11:39:57
 */
package wechat

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"

	"libra/models/constants"
	"libra/pkg/httpUtils"
	"libra/pkg/mapper"
	"libra/pkg/redisUtils"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const (
	// baseUrl        = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"
	// userUrl        = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"
	codeToTokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	tokenUrl       = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	userInfoUrl    = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
	jsTicketUrl    = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

func get(url string, v interface{}) (err error) {
	maps := make(map[string]interface{})
	err = httpUtils.FetchingJson(url, "GET", &maps)
	log.Infof("url: %s; res: %s", url, maps)

	if err != nil {
		return err
	}

	if errcode, ok := maps["errcode"]; ok && errcode != 0 {
		errmsg := maps["errmsg"]
		return errors.New(fmt.Sprintf("%v: %s", errcode, errmsg))
	}

	err = mapper.MapTo(maps, v)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) CodeToToken(code string) (token *WxToken, err error) {
	url := fmt.Sprintf(codeToTokenUrl, client.AppId, client.AppSecret, code)
	token = &WxToken{}
	err = get(url, token)

	return token, nil
}

func (client *Client) GetToken() (token *WxToken, err error) {
	token = &WxToken{}
	key := fmt.Sprintf(constants.RedisWxTokenF, client.AppId)
	json, err := redisUtils.Instance().Get(key).Result()
	if err != redis.Nil {
		mapper.ToObject(json, token)
		return token, nil
	}

	url := fmt.Sprintf(tokenUrl, client.AppId, client.AppSecret)
	err = get(url, token)
	if err != nil {
		return nil, err
	}
	redisUtils.Set(key, token, 7200)
	return token, nil
}

func (client *Client) GetUserInfo(openId string) (wxUserInfo *WxUserInfo, err error) {
	tokenInfo, err := client.GetToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(userInfoUrl, tokenInfo.AccessToken, openId)
	wxUserInfo = &WxUserInfo{}
	err = get(url, wxUserInfo)
	if err != nil {
		return nil, err
	}

	return wxUserInfo, nil
}

func (client *Client) GetJsTicket() (string, error) {
	ticket := ""
	key := fmt.Sprintf(constants.RedisWxJsTicketF, client.AppId)
	ticket, err := redisUtils.Instance().Get(key).Result()
	if err != redis.Nil {
		return ticket, nil
	}

	tokenInfo, err := client.GetToken()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf(jsTicketUrl, tokenInfo.AccessToken)
	maps := make(map[string]interface{})
	err = get(url, &maps)
	if err != nil {
		return "", err
	}
	ticket = maps["ticket"].(string)
	if len(ticket) > 0 {
		redisUtils.Set(key, ticket, 7200)
	}

	return ticket, nil
}

func (client *Client) GetJsSignature(noncestr, timestamp, url string) (string, error) {
	ticket, err := client.GetJsTicket()
	if err != nil {
		return "", err
	}

	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, noncestr, timestamp, url)
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return hex.EncodeToString(sha1.Sum([]byte(""))), nil
}
