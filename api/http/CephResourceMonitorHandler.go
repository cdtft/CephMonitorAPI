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
	conn.ReadDefaultConfigFile()
	conn.Connect()

	ioctx, err := conn.OpenIOContext("k8s")
	if err != nil {
		log.Println("open io context error")
		return
	}
	img := rbd.GetImage(ioctx, imageName)
	log.Println(img)
	ioctx.Destroy()
	conn.Shutdown()
}
