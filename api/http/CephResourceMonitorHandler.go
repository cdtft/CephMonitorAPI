package http

import (
	"CephMonitorAPI/goceph/rados"
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Welcome!")
}

func GetImageUsageByName(response http.ResponseWriter, request *http.Request) {
	//vars := mux.Vars(request)
	//imageName := vars["imageName"]

	conn, _ := rados.NewConn()
	poolNames, _ := conn.ListPools()
	log.Println(poolNames)
}

//func TodoIndex(w http.ResponseWriter, r *http.Request) {
//	todos := Todos{
//		Todo{Name: "Write presentation"},
//		Todo{Name: "Host meetup"},
//	}
//
//	if err := json.NewEncoder(w).Encode(todos); err != nil {
//		panic(err)
//	}
//}

//func TodoShow(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	todoId := vars["todoId"]
//	fmt.Fprintln(w, "Todo show:", todoId)
//}
