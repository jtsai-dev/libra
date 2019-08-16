/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:38:37
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-12 17:20:15
 */
package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": "hello",
	})
}

func View(context *gin.Context) {
	action := strings.TrimLeft(context.Param("action"), "/")
	if len(action) == 0 {
		action = "index"
	}
	context.HTML(http.StatusOK, action+".html", gin.H{
		"title": action,
	})
}
