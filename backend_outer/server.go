package main

import (
	"log"
	"net/http"

	"github.com/ChrisLi03/WatchDOG/backend_outer/management"
)

func handleRequests() {
	http.HandleFunc("/powercubicle/v1/decrpt", management.DecrptCode)
	err := http.ListenAndServe(":12076", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
