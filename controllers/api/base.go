/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:29:21
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-12 16:55:35
 */
package api

import (
	"libra/models"
	"libra/pkg"

	"github.com/gin-gonic/gin"
)

func WJson(c *gin.Context, v interface{}) {
	result := models.R(&v)
	c.JSON(200, result)
}

func WJsonCode(c *gin.Context, code int) {
	result := models.RC(code)
	c.JSON(200, result)
}

func WJsonCodeMsg(c *gin.Context, code int, message string) {
	result := models.RCM(code, message)
	c.JSON(200, result)
}

func GetPageInfo(c *gin.Context) (pageIndex, pageSize int) {
	var pageInfo models.PaginationIn
	err := c.Bind(&pageInfo)
	if err != nil {
		pageInfo.PageIndex = 1
		pageInfo.PageIndex = pkg.Configs.App.PageSize
	}
	if pageInfo.PageIndex == 0 {
		pageInfo.PageIndex = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = pkg.Configs.App.PageSize
	}
	return pageInfo.PageIndex, pageInfo.PageSize
}
