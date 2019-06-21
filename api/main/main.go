package main

import (
	. "CephMonitorAPI/api/http"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":7070", router))
}
