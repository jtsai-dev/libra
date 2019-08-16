/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 15:01:48
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-13 18:01:13
 */
package models

import (
	"time"
)

type Login struct {
	Id          int64
	WxAccountId int64
	AccountId   int64
	Token       string
	Created     time.Time
}
