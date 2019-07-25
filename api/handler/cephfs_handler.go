package handler

import (
	"CephMonitorAPI/goceph/cephfs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func TestCephFS(c *gin.Context)  {
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