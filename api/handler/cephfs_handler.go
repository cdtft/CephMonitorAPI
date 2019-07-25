package handler

import (
	"CephMonitorAPI/goceph/cephfs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

func CreateCephfsDir(c *gin.Context) {
	mount, err := cephfs.CreateMount()
	err = mount.ReadDefaultConfigFile()
	err = mount.Mount()
	if err != nil {
		fmt.Print("create mount error")
	}
	currentDir := mount.CurrentDir()
	log.Println(currentDir)
	err = mount.MakeDir("/k8s/wangcheng", 0755)
	if err != nil {
		log.Println(err.Error())
		return
	}
	mount.ChangeDir("/k8s/wangcheng")
	currentDir = mount.CurrentDir()
	log.Println(currentDir)
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