/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:21:12
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-15 11:34:23
 */
package middlewares

import (
	"bytes"
	"strings"

	"libra/pkg/jsonUtils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// LogHandle log the request and response
func LogHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		if c.Writer.Status() != 404 &&
			strings.Index(c.Request.URL.Path, "api") > -1 {
			res := blw.body.String()

			var in map[string]interface{}
			c.ShouldBindBodyWith(&in, binding.JSON)
			if len(in) < 1 {
				c.ShouldBindBodyWith(&in, binding.XML)
			}
			if _, ok := in["password"]; ok {
				in["password"] = "***"
			}

			str, _ := jsonUtils.ToJson(in)

			log.Infof("%s %s%s; from: %s; request: %s|%s; response: %s",
				c.Request.Method, c.Request.Host, c.Request.URL,
				c.ClientIP(),
				c.Request.Header.Get("Token"), str,
				res)
		}
	}
}
