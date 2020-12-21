package seat

import (
	"fmt"
	"net/http"
)

func SeatRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WatchDOG")
	fmt.Println("Endpoint Hit: SeatRegister")
}
