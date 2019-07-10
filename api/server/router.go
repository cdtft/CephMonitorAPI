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
	rbdApi := router.Group("/api/v1/ceph/rbd")
	{
		rbdApi.POST("/:pool/image/:name/:size", handler.CreateImage)
		rbdApi.DELETE("/:pool/image/:name", handler.DeleteImage)
		rbdApi.GET("/:pool/image/:name/usage", handler.GetImageUsage)
		rbdApi.PUT("/:pool/image/:name/:size", handler.UpdateImageSize)
	}

	//cephfs
	fsApi := router.Group("/api/v1/ceph/fs")
	{
		fsApi.POST("")
	}

	return router
}
