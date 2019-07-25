package handler

import (
	"CephMonitorAPI/api/serializer"
	"CephMonitorAPI/api/service"
	"github.com/gin-gonic/gin"
)

const (
	entries_commond  = "ceph.dir.entries"
	files_commond    = "ceph.dir.files"
	rebytes_commond  = "ceph.dir.rbytes" //目录暂用的字节数
	rctime_commond   = "ceph.dir.rctime"
	rentries_commond = "ceph.dir.rentries"
	rfiles_commond   = "ceph.dir.rfiles"   //目录下的文件数量
	rsubdirs_commond = "ceph.dir.rsubdirs" //子目录个数
	subdirs_commond  = "ceph.dir.subdirs"
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

func DeleteCephDir(c *gin.Context) {

}

func GetCephDirUsage(c *gin.Context) {

}

func GetCephDirInfo(c *gin.Context) {

}

func GetCephDirsInfo(c *gin.Context) {

}

func ChomdCephDir(c *gin.Context) {

}