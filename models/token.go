/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:37:06
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-14 17:48:01
 */
package models

import "time"

type TokenInfo struct {
	Id       int       `json:"accountId"`
	Name     string    `json:"name"`
	Token    string    `json:"token"`
	Expireat time.Time `json:"expireat"`
}

type TokenOut struct {
	NickName string    `json:"nickName"`
	Avatar   string    `json:"avatar"`
	Token    string    `json:"token"`
	Expireat time.Time `json:"expireat"`
	Expirein int       `json:"expirein"`
}
