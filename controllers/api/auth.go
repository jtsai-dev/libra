/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:38:37
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 18:20:39
 */
package api

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"libra/models"
	"libra/models/constants"
	"libra/pkg/conf"
	"libra/pkg/enums"
	"libra/pkg/random"
	"libra/pkg/redisUtils"
	"libra/pkg/wechat"

	"github.com/gin-gonic/gin"
)

// @Summary get WxSession
// @Tags session
// @Produce json
// Param code query {string} true "code"
// @Success 200 {object} models.TokenOut
// @Router /api/wxseession [get]
func WxSession_Get(context *gin.Context) {
	code := context.Query("code")
	if len(code) < 1 {
		panic(enums.ParamsInvalid)
	}

	wxClient := &wechat.Client{
		AppId:     conf.Configs.Wechat.AppId,
		AppSecret: conf.Configs.Wechat.AppSecret,
	}
	token, err := wxClient.CodeToToken(code)
	if err != nil {
		if token != nil {
			WJsonCodeMsg(context, token.Errcode, token.Errmsg)
			return
		}

		panic(err)
	}

	userInfo, err := wxClient.GetUserInfo(token.OpenId)
	if err != nil {
		if userInfo != nil {
			WJsonCodeMsg(context, userInfo.Errcode, userInfo.Errmsg)
			return
		}

		panic(err)
	}

	wxAccount := &models.WxAccount{OpenId: userInfo.OpenId, AppId: conf.Configs.Wechat.AppId}
	has, err := models.X.Get(wxAccount)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	randStr := random.String(32)
	if !has {
		// insert to db
		wxAccount.UnionId = userInfo.UnionId
		wxAccount.NickName = userInfo.Nickname
		wxAccount.Avatar = userInfo.HeadImgurl
		wxAccount.Gender = userInfo.Sex
		wxAccount.Province = userInfo.Province
		wxAccount.City = userInfo.City
		wxAccount.Country = userInfo.Country
		wxAccount.SubscribeStatus = userInfo.Subscribe
		wxAccount.Subscribed = time.Unix(userInfo.SubscribeTime, 0)
		wxAccount.Created = now

		if len(wxAccount.Avatar) < 1 {
			wxAccount.Avatar = conf.Configs.Wechat.DefaultAvatar
		}

		id, _ := models.X.Insert(wxAccount)
		wxAccount.Id = id
	}

	// token's format: id.name.key
	key := fmt.Sprintf("%d.%s.%s", wxAccount.Id, wxAccount.NickName, randStr)
	tokenStr := base64.StdEncoding.EncodeToString([]byte(key))
	out := models.TokenOut{
		Token:    tokenStr,
		Expirein: conf.Configs.App.TokenExpiredSeconds,
		Expireat: now.Add(time.Duration(conf.Configs.App.TokenExpiredSeconds) * time.Second),
	}

	// delay 30s for network
	accountTokenKey := fmt.Sprintf(constants.RedisAccountTokenF, randStr)
	redisUtils.Set(accountTokenKey, wxAccount, conf.Configs.App.TokenExpiredSeconds+30)

	login := &models.Login{
		WxAccountId: wxAccount.Id,
		Token:       tokenStr,
		Created:     now,
	}
	go logLogin(login)

	WJson(context, out)
}

func logLogin(login *models.Login) {
	last := &models.Login{}
	has, _ := models.X.Where("wx_account_id = ? AND created >= ?",
		login.WxAccountId,
		login.Created.Add(-time.Duration(conf.Configs.App.TokenExpiredSeconds)*time.Second)).
		Desc("id", "created").
		Get(last)
	if has {
		bytes, _ := base64.StdEncoding.DecodeString(last.Token)
		key := strings.Split(string(bytes), ".")[2]

		tokenKey := fmt.Sprintf(constants.RedisAccountTokenF, key)
		redisUtils.Instance().Del(tokenKey)
	}
	models.X.Insert(login)
}
