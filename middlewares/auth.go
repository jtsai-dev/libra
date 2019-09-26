/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:40:56
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-15 11:05:30
 */
package middlewares

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"libra/models"
	"libra/models/constants"
	"libra/pkg/enums"
	"libra/pkg/mapper"
	"libra/pkg/redisUtils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// AuthHandle would valid the token in header
func AuthHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		if len(token) == 0 {
			unAuthHandle(c)
		}

		bytes, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			unAuthHandle(c)
		}

		arr := strings.Split(string(bytes), ".")
		if len(arr) < 3 {
			unAuthHandle(c)
		}
		id, _ := strconv.ParseInt(arr[0], 10, 64)
		// name := arr[1]
		key := arr[2]

		tokenKey := fmt.Sprintf(constants.RedisAccountTokenF, key)
		value, err := redisUtils.Instance().Get(tokenKey).Result()
		if err == redis.Nil {
			unAuthHandle(c)
		}
		var wxAccount models.WxAccount
		mapper.ToObject(value, &wxAccount)

		if wxAccount.Id != id {
			unAuthHandle(c)
		}

		c.Set(constants.SessionAccount, wxAccount)

		c.Next()
	}
}

func unAuthHandle(c *gin.Context) {
	panic(enums.AuthTokenInvalid)
}
