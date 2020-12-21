package main

import (
	"log"
	"net/http"

	"github.com/ChrisLi03/WatchDOG/backend_outer/seat"
	"github.com/ChrisLi03/WatchDOG/backend_outer/user"
)

func handleRequests() {
	http.HandleFunc("/powercubicle/v1/seat", seat.RetrieveAllSeatStatus)
	http.HandleFunc("/powercubicle/v1/seat/register", seat.SeatRegister)
	http.HandleFunc("/powercubicle/v1/user/login", user.UserLogin)
	http.HandleFunc("/powercubicle/v1/user/register", user.UserRegister)

	err := http.ListenAndServe(":12077", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
