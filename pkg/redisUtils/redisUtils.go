/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:22:56
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-12 14:23:39
 */
package redisUtils

import (
	"encoding/json"
	"time"

	"libra/pkg"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var client *redis.Client

func Instance() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     pkg.Configs.Redis.Address,
			Password: pkg.Configs.Redis.Password,
			DB:       pkg.Configs.Redis.DBIndex,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Error(err)
			panic("warning, redis connection is not alive")
		}
	}

	return client
}

func Set(key string, value interface{}, expiredSecods int) *redis.StatusCmd {
	bytes, _ := json.Marshal(value)
	return Instance().Set(
		key,
		bytes,
		time.Duration(expiredSecods)*time.Second)
}

func HSet(key, field string, value interface{}) *redis.BoolCmd {
	bytes, _ := json.Marshal(value)
	return Instance().HSet(
		key,
		field,
		bytes)
}
