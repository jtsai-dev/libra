/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-02 10:22:05
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 18:17:29
 */
package routers

import (
	"libra/controllers"
	"libra/controllers/api"
	v1 "libra/controllers/api/v1"
	"libra/middlewares"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Regist(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery())

	router.Use(
		middlewares.LogHandle(),
		middlewares.CORSHandle(),
		middlewares.GlobalErrorHandle())

	// views & static
	router.GET("/", controllers.Index)
	router.GET("view/*action", controllers.View)
	router.Static("static", "static")
	router.LoadHTMLGlob("templates/*")

	// api
	apiGroup := router.Group("api")
	{
		// wechat
		apiGroup.GET("wechat", api.Wechat_Get)
		apiGroup.POST("wechat", api.Wechat_Post)

		// session
		apiGroup.POST("wxsession", api.WxSession_Post)

		v1Group := apiGroup.Group("v1", middlewares.AuthHandle())
		{
			// adjudication
			v1Group.GET("adjudication/history", v1.History_Get)
			v1Group.GET("adjudication", v1.Adjudication_Get)

			// node
			v1Group.GET("node", v1.Node_Get)
			v1Group.POST("node", v1.Node_Post)
			v1Group.PUT("node", v1.Node_Put)
			v1Group.DELETE("node", v1.Node_Delete)
		}
	}

	// swagger
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
