package server

import (
	"CephMonitorAPI/api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()

	//应用健康检测
	v1 := router.Group("/api/v1/")
	{
		v1.POST("ping", handler.Ping)
	}

	//rbd
	rbdApi := router.Group("/api/v1/ceph/:pool/rbd")
	{
		rbdApi.GET("image/:name", handler.GetImageInfo)
		rbdApi.POST("image", handler.CreateImage)
		rbdApi.DELETE("image", handler.DeleteImage)
		rbdApi.GET("image/usage", handler.GetImageUsage)
		rbdApi.PUT("image/size", handler.UpdateImageSize)
	}

	//cephfs
	fsApi := router.Group("/api/v1/ceph/fs")
	{
		fsApi.POST("")
	}

	return router
}
