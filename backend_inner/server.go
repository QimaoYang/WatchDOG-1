package main

import (
	"log"
	"net/http"

	"github.com/ChrisLi03/WatchDOG/backend_inner/management"
)

func handleRequests() {
	http.HandleFunc("/powercubicle/v1/seat/encrypt", management.EncryptCode)
	err := http.ListenAndServe(":12076", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
