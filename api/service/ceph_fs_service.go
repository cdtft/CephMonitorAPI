package service

import (
	"CephMonitorAPI/goceph/cephfs"
	"log"
)

type CephfsService struct {
	Dir string
}

func (cephfsService *CephfsService) CreateDir() error {
	mount, err := cephfs.CreateMount()
	err = mount.ReadDefaultConfigFile()
	err = mount.Mount()
	if err != nil {
		log.Println("create mount error")
		return err
	}
	currentDir := mount.CurrentDir()
	log.Println(currentDir)
	err = mount.MakeDir("/k8s/wangcheng", 0755)
	if err != nil {
		log.Printf("Make dir:%s error", cephfsService.Dir)
		return err
	}
	return nil
}