package user

import (
	"fmt"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WatchDOG")
	fmt.Println("Endpoint Hit: UserLogin")
}
