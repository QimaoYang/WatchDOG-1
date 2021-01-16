package main

import (
	"log"
	"net/http"

	"github.com/ChrisLi03/WatchDOG/backend_outer/seat"
	"github.com/ChrisLi03/WatchDOG/backend_outer/user"
	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter()
	wd_router := router.PathPrefix("/powercubicle/v1").Subrouter()
	wd_router.HandleFunc("/seat", seat.RetrieveAllSeatStatus).Methods("GET")
	wd_router.HandleFunc("/seat/register", seat.SeatRegister).Methods("POST")
	wd_router.HandleFunc("/seat/release", seat.SeatRelease).Methods("POST")
	wd_router.HandleFunc("/user/login", user.UserLogin).Methods("POST")
	wd_router.HandleFunc("/user/register", user.UserRegister).Methods("POST")
	wd_router.HandleFunc("/user/seat", user.UserSeat).Methods("GET")

	err := http.ListenAndServe(":12077", wd_router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
