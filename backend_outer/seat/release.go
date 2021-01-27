package seat

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ChrisLi03/WatchDOG/backend_outer/common"
)

func SeatRelease(w http.ResponseWriter, r *http.Request) {
	seatSessionKey := r.Header.Get("Authorization")
	releaseSeat(w, r, seatSessionKey)
}

func releaseSeat(w http.ResponseWriter, r *http.Request, sessionAuth string) {
	urlSeatRelease := "http://localhost:5001/powercubicle/v1/db/seat/release"
	log.Println("[WD] Start releasing your seat")

	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, err := http.NewRequest(http.MethodPost, urlSeatRelease, nil)
	req.Header.Add("Authorization", sessionAuth)

	if err != nil {
		log.Fatal(err)
	}

	resp, getErr := cubeClient.Do(req)

	switch resp.StatusCode {
	case 400:
		errMsg := "Don't double release the seat"
		resError := common.Errors{}
		resError = resError.NewError(400, errMsg)
		errCode, errMsg := resError.GetError()

		log.Printf("[WD] ", errMsg)
		http.Error(w, errMsg, errCode)
		return
	case 401:
		errMsg := "User's token has expired"
		resError := common.Errors{}
		resError = resError.NewError(401, errMsg)
		errCode, errMsg := resError.GetError()

		log.Printf("[WD] ", errMsg)
		http.Error(w, errMsg, errCode)
		return
	}

	if getErr != nil {
		log.Fatal(getErr)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Your seat has been release")
}
