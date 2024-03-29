/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 10:34:32
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 17:39:01
 */
package main

import (
	"fmt"
	"math/rand"
	"time"

	"libra/models"
	"libra/pkg/conf"
	"libra/pkg/logger"
	"libra/routers"

	_ "libra/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/gin-gonic/gin"
)

func init() {
	go rand.Seed(time.Now().UnixNano())
	conf.Setup()
	logger.Setup(conf.Configs.Log.Path)
	models.Setup()
	// models.SyncDataBase()
}

func main() {
	gin.SetMode(conf.Configs.Server.RunMode)

	router := gin.Default()
	routers.Regist(router)

	router.Run(fmt.Sprintf(":%d", conf.Configs.Server.Port))
}
