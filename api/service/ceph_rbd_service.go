package service

import (
	"CephMonitorAPI/goceph/rados"
	"CephMonitorAPI/goceph/rbd"
	"errors"
	"log"
)

type Image struct {
	Pool string `uri:"pool" json:"pool"`
	Name string `uri:"name" json:"name"`
	Size uint64 `uri:"size" json:"size"`
}

func (image *Image) Create() error {
	log.Printf("创建云盘pool:%s, name:%s, size:%d \n", image.Pool, image.Name, image.Size)
	conn, _ := rados.NewConn()
	conn.ReadDefaultConfigFile()
	conn.Connect()
	ioctx, err := conn.OpenIOContext(image.Pool)

	if err != nil {
		return errors.New("云盘创建失败:" + err.Error())
	}
	_, err = rbd.Create(ioctx, image.Name, image.Size*1024*1024*1024, 22)
	if err != nil {
		return errors.New("云盘创建失败:" + err.Error())
	}
	ioctx.Destroy()
	conn.Shutdown()
	return nil
}

func (image *Image) Delete() error {
	log.Printf("删除云盘pool:%s, name:%s \n", image.Pool, image.Name)
	conn, _ := rados.NewConn()
	
	return nil
}
