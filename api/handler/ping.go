package handler

import (
	"CephMonitorAPI/api/serializer"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, serializer.ResponseJSON{
		Code: 1200,
		Msg:  "i am OK",
	})
}
