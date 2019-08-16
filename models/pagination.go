/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:33:46
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-12 15:27:08
 */
package models

import "math"

type PaginationIn struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

type PaginationOut struct {
	Items      interface{} `json:"items"`
	PageIndex  int         `json:"pageIndex"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int64       `json:"totalPage"`
	TotalCount int64       `json:"totalCount"`
}

func (p *PaginationOut) Set(pageIndex int, pageSize int, totalCount int64) *PaginationOut {
	p.PageIndex = pageIndex
	p.PageSize = pageSize
	p.TotalCount = totalCount

	totalCountF := float64(totalCount)
	pageSizeF := float64(pageSize)
	p.TotalPage = (int64)(math.Ceil(totalCountF / pageSizeF))

	return p
}
