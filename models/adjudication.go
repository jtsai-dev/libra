/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-13 17:49:05
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 16:41:33
 */
package models

import "time"

type Adjudication struct {
	Id            int64
	WxAccountId   int64
	DirectoryId   int64
	DirectoryName string
	OptionId      int64
	OptionName    string

	Created time.Time
}

type AdjudicationOut struct {
	DirectoryId   int64
	DirectoryName string
	OptionId      int64
	OptionName    string
}
