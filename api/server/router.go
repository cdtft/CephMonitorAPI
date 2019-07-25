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

	poolApi := router.Group("/api/v1/ceph/pool")
	{
		poolApi.POST("/:poolName", handler.CreatePool)
		poolApi.DELETE("/:poolName", handler.DeletePool)
	}

	//rbd
	rbdApi := router.Group("/api/v1/ceph/rbd")
	{
		rbdApi.POST("/:pool/image/:name/:size", handler.CreateImage)
		rbdApi.DELETE("/:pool/image/:name", handler.DeleteImage)
		rbdApi.GET("/:pool/image/:name/usage", handler.GetImageUsage)
		rbdApi.PUT("/:pool/image/:name/:size", handler.UpdateImageSize)

		//batch
		rbdApi.GET("/:pool/images/usages", handler.GetImagesInfos)
		rbdApi.POST("/:pool/images", handler.CreateImages)
		rbdApi.DELETE("/:pool/images", handler.DeleteImages)
	}

	//cephfs
	fsApi := router.Group("/api/v1/ceph/fs")
	{
		fsApi.POST("/test", handler.TestCephFS)
	}

	return router
}
