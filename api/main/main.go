package main

import (
	"CephMonitorAPI/api/server"
)

func main() {
	router := server.NewRouter()
	router.Run(":10086")
}
