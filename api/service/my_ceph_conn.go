package service

import (
	"CephMonitorAPI/goceph/rados"
	"log"
)

var myCephConn *rados.Conn

func init() {
	log.Print("初始化ceph链接")
	//InstanceMyCephConn()
}

//获得连接
func InstanceMyCephConn() *rados.Conn {
	if myCephConn == nil {
		myCephConn, _ = rados.NewConn()
		myCephConn.ReadDefaultConfigFile()
		myCephConn.Connect()
	}
	return myCephConn
}

//关闭连接
func CloseMyCephConn() {
	if myCephConn != nil {
		myCephConn.Shutdown()
	}
}
