package main

import (
	"CephMonitorAPI/api/server"
	"log"
)

func main() {
	router := server.NewRouter()
 	err := router.Run("0.0.0.0:10086")
 	if err != nil {
		log.Println("应用启动失败", err)
	}
 	log.Println("CephMonitorAPI启动成功")
}
