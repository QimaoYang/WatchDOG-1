package user

import (
	"fmt"
	"net/http"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WatchDOG")
	fmt.Println("Endpoint Hit: UserRegister")
}
