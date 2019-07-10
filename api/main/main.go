package main

import (
	"CephMonitorAPI/api/server"
	"log"
)

func main() {
	router := server.NewRouter()
 	err := router.Run(":10086")
 	if err != nil {
		log.Println("应用启动失败", err)
	}
}
