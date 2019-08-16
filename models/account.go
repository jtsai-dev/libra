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

type Account struct {
	Id       int64
	Name     string
	IdNumber string
	Mobile   string
	Email    string
	Password string
	Remark   string
	Status   int
	Created  time.Time
	Updated  time.Time
	Deleted  time.Time
}
