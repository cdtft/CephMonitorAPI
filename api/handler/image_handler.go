package handler

import (
	"CephMonitorAPI/api/service"
	"log"
	"github.com/gin-gonic/gin"
)

// 获取image信息
func GetImageInfo(ctx *gin.Context) {
	var image service.Image
	if err := ctx.ShouldBindWith(&image, ); err == nil {

	}
	log.Printf("")
}

// 创建image
func CreateImage(ctx *gin.Context) {

}

// 删除image
func DeleteImage(ctx *gin.Context) {

}

// 获取image使用率
func GetImageUsage(ctx *gin.Context) {

}

// resize
func UpdateImageSize(ctx *gin.Context) {

}
