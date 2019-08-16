/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:14:05
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-12 14:35:54
 */
package middlewares

import (
	"libra/models"
	"libra/pkg"
	"libra/pkg/enums"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type FriendlyError struct {
	error
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// FriendlyErrorHandle only handle FriendlyError, use it after gin.Recovery if you use
func GlobalErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var result *models.Result
				if fe, ok := err.(FriendlyError); ok {
					if fe.Code == 0 {
						fe.Code = enums.Fail
					}
					if len(fe.Message) == 0 {
						fe.Message = enums.GetRespCodeDesc(fe.Code)
					}
					result = models.RCM(fe.Code, fe.Message)
				} else if str, ok := err.(string); ok {
					result = models.RCM(enums.Fail, str)
				} else if code, ok := err.(int); ok {
					result = models.RC(code)
				} else if pkg.Configs.Server.RunMode == "debug" {
					if e, ok := err.(error); ok {
						log.Error("Panic: %v", err)
						result = models.RCM(enums.Fail, e.Error())
					}
				} else {
					result = models.RC(enums.Fail)
				}
				if result != nil {
					c.JSON(200, result)
					c.Abort()
					return
				}
			}
		}()

		c.Next()
	}
}
