package server

import (
	"CephMonitorAPI/api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("/api/v1/ceph")
	{
		v1.POST("ping", handler.Ping)
	}
	return router
}
