package seat

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func SeatRelease(w http.ResponseWriter, r *http.Request) {
	seatSessionKey := r.Header.Get("Auth")
	log.Println("the seat session key is", seatSessionKey)
	fmt.Println("Endpoint Hit: SeatRegister")
	releaseSeat(w, r, seatSessionKey)
}

func releaseSeat(w http.ResponseWriter, r *http.Request, sessionAuth string) {
	urlSeatRelease := "http://localhost:5001/powercubicle/v1/seat/release"

	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	fmt.Println("auth", sessionAuth)
	req, err := http.NewRequest(http.MethodPost, urlSeatRelease, nil)
	req.Header.Add("Authorization", sessionAuth)

	if err != nil {
		log.Fatal(err)
	}

	_, getErr := cubeClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}
}
