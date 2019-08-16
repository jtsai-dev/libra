/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 15:01:48
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-13 18:01:20
 */
package models

import (
	"time"
)

type WxAccount struct {
	Id              int64
	AccountId       int64
	UnionId         string
	AppId           string
	OpenId          string // 用户的唯一标识
	NickName        string // 用户昵称
	Avatar          string // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效
	Gender          int    // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province        string // 用户个人资料填写的省份
	City            string // 普通用户个人资料填写的城市
	Country         string // 国家，如中国为CN
	SubscribeStatus int
	Subscribed      time.Time
	UnSubscribed    time.Time
	Created         time.Time
	Updated         time.Time
	Deleted         time.Time
}
