package router

import (
	v1 "geek/04/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 创建v1组
	vg1 := router.Group("/v1")
	{
		// 在v1这个分组下，注册路由
		vg1.GET("/user/:id", v1.GetUserHttp)
		vg1.POST("/user/:id", v1.PostUserHttp)
	}

	return router
}
