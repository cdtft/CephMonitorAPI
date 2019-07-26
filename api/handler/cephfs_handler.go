package handler

import (
	"CephMonitorAPI/api/serializer"
	"CephMonitorAPI/api/service"
	"github.com/gin-gonic/gin"
)

func CreateCephfsDir(ctx *gin.Context) {
	var cephfsService service.CephfsService
	if err := ctx.ShouldBindUri(&cephfsService); err == nil {
		if err := cephfsService.CreateDir(); err == nil {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1200,
				Msg:  "创建成功",
				Data: nil,
			})
		} else {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1500,
				Msg:  err.Error(),
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

func DeleteCephDir(ctx *gin.Context) {
	var cephfsService service.CephfsService
	if err := ctx.ShouldBindUri(&cephfsService); err == nil {
		if err := cephfsService.DeleteDir(); err == nil {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1200,
				Msg:  "删除成功",
				Data: nil,
			})
		} else {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1500,
				Msg:  err.Error(),
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

func GetCephDirUsage(ctx *gin.Context) {
	var cephfsService service.CephfsService
	if err := ctx.ShouldBindUri(&cephfsService); err == nil {
		if cephfsDir, err := cephfsService.GetDirUsage(); err == nil {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1200,
				Msg:  "查询成功",
				Data: cephfsDir,
			})
		} else {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1500,
				Msg:  err.Error(),
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

func GetCephDirsUsage(ctx *gin.Context) {
	var cephfsDirBatchService service.CephfsDirBatchService
	if err := ctx.ShouldBindUri(&cephfsDirBatchService); err == nil {
		if dirUsageArray, err := cephfsDirBatchService.GetCephDirsUsage(); err == nil {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1200,
				Msg:  "查询成功",
				Data: dirUsageArray,
			})
		} else {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1500,
				Msg:  err.Error(),
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

func ChomdCephDir(ctx *gin.Context) {
	var cephfsService service.CephfsService
	if err := ctx.ShouldBindUri(&cephfsService); err == nil {
		if err := cephfsService.DeleteDir(); err == nil {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1200,
				Msg:  "删除成功",
				Data: nil,
			})
		} else {
			ctx.JSON(200, serializer.ResponseJSON{
				Code: 1500,
				Msg:  err.Error(),
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