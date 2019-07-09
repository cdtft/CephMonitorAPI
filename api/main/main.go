package main

import (
	myhttp "CephMonitorAPI/api/http"
	"log"
	"net/http"
)

func main() {
	router := myhttp.NewRouter()
	log.Println("api server is start listener on 7070 port")
	log.Fatal(http.ListenAndServe(":7070", router))
}
