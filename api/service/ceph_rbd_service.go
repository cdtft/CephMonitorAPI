package service

import (
	"CephMonitorAPI/goceph/rados"
	"CephMonitorAPI/goceph/rbd"
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

type ImageService struct {
	Pool string `uri:"pool" json:"pool"`
	Name string `uri:"name" json:"name"`
	Size uint64 `uri:"size" json:"size"`
}

type PoolService struct {
	Name string `uri:"name" json:"name"`
}

func (image *ImageService) Create() error {
	log.Printf("创建云盘pool:%s, name:%s, size:%d \n", image.Pool, image.Name, image.Size)
	conn, _ := rados.NewConn()
	_ = conn.ReadDefaultConfigFile()
	_ = conn.Connect()
	ioctx, err := conn.OpenIOContext(image.Pool)

	if err != nil {
		return errors.New("云盘创建失败:" + err.Error())
	}
	if _, err = rbd.Create(ioctx, image.Name, image.Size*1024*1024*1024, 22); err != nil {
		return errors.New("云盘创建失败:" + err.Error())
	}
	ioctx.Destroy()
	conn.Shutdown()
	return nil
}

func (image *ImageService) Delete() error {
	log.Printf("删除云盘pool:%s, name:%s \n", image.Pool, image.Name)
	conn, _ := rados.NewConn()
	defer conn.Shutdown()
	_ = conn.ReadDefaultConfigFile()
	_ = conn.Connect()
	ioctx, err := conn.OpenIOContext(image.Pool)
	defer ioctx.Destroy()
	if err != nil {
		return errors.New("云盘删除失败:" + err.Error())
	}
	cephImage := rbd.GetImage(ioctx, image.Name)
	if err = cephImage.Remove(); err != nil {
		return errors.New("云盘删除失败:" + err.Error())
	}
	return nil
}

func (image *ImageService) Resize() error {
	log.Printf("修改云盘大小pool:%s, name:%s, size:%d", image.Pool, image.Name, image.Size)
	conn, _ := rados.NewConn()
	defer conn.Shutdown()
	_ = conn.ReadDefaultConfigFile()
	_ = conn.Connect()
	ioctx, err := conn.OpenIOContext(image.Pool)
	defer ioctx.Destroy()
	if err != nil {
		return errors.New("云盘删除失败:" + err.Error())
	}
	cephImage := rbd.GetImage(ioctx, image.Name)
	originSize, _ := cephImage.GetSize()
	resize := image.Size * 1024 * 1024 * 102
	if originSize > resize {
		return errors.New("修改云盘的大小，小于原大小")
	}
	if err = cephImage.Resize(resize); err != nil {
		return errors.New("云盘删除失败:" + err.Error())
	}
	return nil
}

func (image *ImageService) GetUsage() (used string, error error) {
	log.Printf("执行统计命令pool:%s, name:%s", image.Pool, image.Name)
	cmd := exec.Command("/bin/bash", "-c", "rbd du "+image.Pool+"/"+image.Name)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "-1", errors.New("云盘统计命令执行失败")
	}
	resultStringArray := strings.Fields(out.String())
	usedStr := resultStringArray[5]
	return usedStr, nil
}

func (pool *PoolService) CreatePool() error {
	log.Printf("创建pool:%s", pool.Name)
	conn, _ := rados.NewConn()
	_ = conn.ReadDefaultConfigFile()
	_ = conn.Connect()
	defer conn.Shutdown()
	if err := conn.MakePool(pool.Name); err != nil {
		return errors.New("创建pool失败！")
	}
	return nil
}

//这是一个很危险的方法
func (pool *PoolService) DeletePool() error {
	log.Printf("创建pool:%s", pool.Name)
	conn, _ := rados.NewConn()
	_ = conn.ReadDefaultConfigFile()
	_ = conn.Connect()
	defer conn.Shutdown()
	if err := conn.DeletePool(pool.Name); err != nil {
		return errors.New("删除pool失败！")
	}
	return nil
}

