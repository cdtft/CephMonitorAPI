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
	if err == nil {
		fmt.Print("create mount error")
	}
	currentDir := mount.CurrentDir()
	log.Print(currentDir)
}