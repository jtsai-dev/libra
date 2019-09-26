/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 16:55:26
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 18:18:41
 */
package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"libra/models"
	"libra/pkg/conf"
	"libra/pkg/enums"
	"libra/pkg/random"
	"libra/pkg/wechat"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary config for wechat's settings
// @Tags wechat
// @Success 200 {string} string
// @Router /api/wechat [get]
func Wechat_Get(context *gin.Context) {
	echostr := context.Query("echostr")
	fmt.Println(echostr)
	context.String(200, echostr)
}

// @Summary Listening WechatEvent
// @Tags wechat
// @Produce xml
// @Success 200 object models.Result
// @Router /api/wechat [post]
func Wechat_Post(context *gin.Context) {

	go func(context *gin.Context) {
		var in wechat.WxEvent

		if err := context.ShouldBindBodyWith(&in, binding.XML); err != nil {

			var bodyBytes []byte
			if context.Request.Body != nil {
				bodyBytes, _ = ioutil.ReadAll(context.Request.Body)
			}

			// Restore the io.ReadCloser to its original state
			context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			fmt.Println(string(bodyBytes))
			log.Println(err)
			return
		}

		switch in.MsgType {
		case "event":
			wxEventHandle(&in)
		}
	}(context)

	context.String(200, "")
}

func wxEventHandle(in *wechat.WxEvent) {
	switch in.Event {
	case "subscribe", "unsubscribe":
		wxSubscribeHandle(in)
	}
}

func wxSubscribeHandle(in *wechat.WxEvent) {
	wxClient := &wechat.Client{
		AppId:     conf.Configs.Wechat.AppId,
		AppSecret: conf.Configs.Wechat.AppSecret,
	}

	has, _ := models.X.Exist(&models.WxAccount{
		OpenId: in.FromUserName,
		AppId:  wxClient.AppId,
	})

	entity := &models.WxAccount{
		AppId:  wxClient.AppId,
		OpenId: in.FromUserName}

	isSync := false
	now := time.Now()
	cols := []string{"subscribe_status", "updated"}
	if in.Event == "subscribe" {
		cols = append(cols, "subscribed")

		entity.Subscribed = now
		entity.SubscribeStatus = enums.Subscribed

		wxUserInfo, err := wxClient.GetUserInfo(entity.OpenId)
		if err == nil {
			isSync = true

			entity.UnionId = wxUserInfo.UnionId
			entity.NickName = wxUserInfo.Nickname
			entity.Gender = wxUserInfo.Sex
			entity.Avatar = wxUserInfo.HeadImgurl
			entity.Country = wxUserInfo.Country
			entity.Province = wxUserInfo.Province
			entity.City = wxUserInfo.City

			if len(entity.Avatar) < 1 {
				entity.Avatar = conf.Configs.Wechat.DefaultAvatar
			}
		}
	} else {
		cols = append(cols, "un_subscribed")

		entity.UnSubscribed = now
		entity.SubscribeStatus = enums.UnSubscribed
	}

	if has {
		entity.Updated = now

		if isSync {
			cols = append(cols,
				"union_id", "nick_name", "gender", "avatar", "country", "province", "city")
		}

		models.X.Cols(cols...).
			Where("app_id = ? AND open_id = ?", entity.AppId, entity.OpenId).
			Update(entity)
	} else {
		entity.Created = now
		id, er := models.X.Insert(entity)
		fmt.Println(id, er)
	}
}

func WechatJsTicket_Get(context *gin.Context) {
	wxClient := &wechat.Client{
		AppId:     conf.Configs.Wechat.AppId,
		AppSecret: conf.Configs.Wechat.AppSecret,
	}

	noncestr := random.String(16)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	var url string
	context.ShouldBindBodyWith(&url, binding.JSON)
	signature, err := wxClient.GetJsSignature(noncestr, timestamp, url)

	if err != nil {
		panic(err)
	}

	WJson(context, signature)
}
