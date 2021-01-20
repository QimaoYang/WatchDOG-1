package main

import (
	"log"
	"net/http"

	"github.com/ChrisLi03/WatchDOG/backend_outer/seat"
	"github.com/ChrisLi03/WatchDOG/backend_outer/user"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter()
	wd_router := router.PathPrefix("/powercubicle/v1").Subrouter()
	wd_router.HandleFunc("/seat", seat.RetrieveAllSeatStatus).Methods("GET", "OPTIONS")
	wd_router.HandleFunc("/seat/register", seat.SeatRegister).Methods("POST", "OPTIONS")
	wd_router.HandleFunc("/seat/release", seat.SeatRelease).Methods("POST", "OPTIONS")
	wd_router.HandleFunc("/user/login", user.UserLogin).Methods("POST", "OPTIONS")
	wd_router.HandleFunc("/user/register", user.UserRegister).Methods("POST", "OPTIONS")
	wd_router.HandleFunc("/user/seat", user.UserSeat).Methods("GET", "OPTIONS")

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	err := http.ListenAndServe(":12077", handlers.CORS(header, methods, origins)(wd_router))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
