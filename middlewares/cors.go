/*
 * @Descripttion: 
 * @version: 
 * @Author: yyeiei
 * @Date: 2019-08-12 14:13:02
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-12 14:36:18
 */
package middlewares

import (
	"github.com/gin-gonic/gin"
)

// CORSHandle allow Cross-origin_resource_sharing
func CORSHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
