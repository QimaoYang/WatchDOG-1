package qr_code

import (
	"fmt"
	"net/http"
)

func EncrptCode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WatchDOG")
	fmt.Println("Endpoint Hit: WatchDOG")
}
