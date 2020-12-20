package main

import (
	"log"
	"net/http"

	"github.com/ChrisLi03/WatchDOG/backend_inner/qr_code"
)

func handleRequests() {
	http.HandleFunc("/powercubicle/v1/qrcode", qr_code.EncrptCode)
	err := http.ListenAndServe(":12076", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
