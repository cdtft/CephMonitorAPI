package service

import (
	"CephMonitorAPI/goceph/rados"
	"log"
)

var MyCephConn *rados.Conn

func init() {
	log.Print("初始化ceph链接")
	InstanceMyCephConn()
}

//获得连接
func InstanceMyCephConn() *rados.Conn {
	if MyCephConn != nil {
		MyCephConn, _ = rados.NewConn()
		_ = MyCephConn.ReadDefaultConfigFile()
		_ = MyCephConn.Connect()
	}
	return MyCephConn
}

//关闭连接
func CloseMyCephConn() {
	if MyCephConn != nil {
		MyCephConn.Shutdown()
	}
}
