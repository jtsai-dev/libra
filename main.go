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
	"libra/pkg"
	"libra/pkg/logUtils"
	"libra/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	go rand.Seed(time.Now().UnixNano())
	pkg.Setup()
	logUtils.Setup()
	models.Setup()
	// models.SyncDataBase()
}

func main() {
	gin.SetMode(pkg.Configs.Server.RunMode)

	router := gin.Default()
	routers.Regist(router)

	router.Run(fmt.Sprintf(":%d", pkg.Configs.Server.Port))
}
