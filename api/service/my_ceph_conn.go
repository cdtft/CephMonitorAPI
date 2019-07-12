package service

import (
	"CephMonitorAPI/goceph/rados"
	"log"
)

var MyCephConn *rados.Conn

func init() {
	log.Print("初始化ceph链接")
	//InstanceMyCephConn()
}

func InstanceMyCephConn() *rados.Conn {
	if MyCephConn != nil {
		MyCephConn, _ = rados.NewConn()
		_ = MyCephConn.ReadDefaultConfigFile()
		_ = MyCephConn.Connect()
	}
	return MyCephConn
}

func CloseMyConn() {
	if MyCephConn != nil {
		MyCephConn.Shutdown()
	}
}
