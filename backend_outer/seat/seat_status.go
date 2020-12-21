package seat

import (
	"fmt"
	"net/http"
)

func RetrieveAllSeatStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WatchDOG")
	fmt.Println("Endpoint Hit: RetrieveAllSeatStatus")
}
