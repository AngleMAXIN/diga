package routers

import (
	v1 "Analysis-statistics/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode("debug")

	apiv1 := r.Group("/api.v2")
	{
		//获取标签列表
		apiv1.GET("/user/:userID/statistics", v1.UserPortraitStatistics)
	}

	return r
}
