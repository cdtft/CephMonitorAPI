package handler

import (
	"CephMonitorAPI/api/serializer"
	"CephMonitorAPI/api/service"

	"github.com/gin-gonic/gin"
)

// 创建image
func CreateImage(ctx *gin.Context) {
	var imageService service.ImageService
	if err := ctx.ShouldBindUri(&imageService); err == nil {
		if err := imageService.Create(); err == nil {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1200,
				Msg:  "创建成功",
				Data: nil,
			})
		} else {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1500,
				Msg: err.Error(),
				Data: nil,
			})
		}
	} else {
		ctx.JSON(400, serializer.ResponseJSON{
			Code: 1400,
			Msg:  "参数绑定失败",
			Data: nil,
		})
	}
}

// 删除image
func DeleteImage(ctx *gin.Context) {
	var imageService service.ImageService
	if err := ctx.ShouldBindUri(&imageService); err == nil {
		if err := imageService.Delete(); err == nil {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1200,
				Msg:  "删除成功",
				Data: nil,
			})
		} else {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1500,
				Msg: err.Error(),
				Data: nil,
			})
		}
	} else {
		ctx.JSON(400, serializer.ResponseJSON{
			Code: 1400,
			Msg:  "参数绑定失败",
			Data: nil,
		})
	}
}

// 获取image使用率
func GetImageUsage(ctx *gin.Context) {
	var imageService service.ImageService
	if err := ctx.ShouldBindUri(&imageService); err == nil {
		if used, err := imageService.GetUsage(); err == nil {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1200,
				Msg:  "获取已使用大小",
				Data: used,
			})
		} else {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1500,
				Msg: err.Error(),
				Data: nil,
			})
		}
	} else {
		ctx.JSON(400, serializer.ResponseJSON{
			Code: 1400,
			Msg:  "参数绑定失败",
			Data: nil,
		})
	}
}

// resize
func UpdateImageSize(ctx *gin.Context) {
	var imageService service.ImageService
	if err := ctx.ShouldBindUri(&imageService); err == nil {
		if err := imageService.Delete(); err == nil {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1200,
				Msg:  "删除成功",
				Data: nil,
			})
		} else {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1500,
				Msg: err.Error(),
				Data: nil,
			})
		}
	} else {
		ctx.JSON(400, serializer.ResponseJSON{
			Code: 1400,
			Msg:  "参数绑定失败",
			Data: nil,
		})
	}
}
