package http

import (
	"CephMonitorAPI/goceph/rados"
	"CephMonitorAPI/goceph/rbd"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Welcome!")
}

func GetImageUsageByName(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	imageName := vars["imageName"]
	conn, _ := rados.NewConn()
	err := conn.ReadDefaultConfigFile()
	if err != nil {
		log.Println("read default config file error!")
		return
	}
	err = conn.Connect()
	if err != nil {
		log.Println("connect error!")
		return
	}
	ctx, err := conn.OpenIOContext("k8s")
	if err != nil {
		log.Println("open io context error!")
		return
	}
	img := rbd.GetImage(ctx, imageName)

	log.Println(img.GetSize())
	ctx.Destroy()
	conn.Shutdown()
}
