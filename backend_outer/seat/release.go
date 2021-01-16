package seat

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ChrisLi03/WatchDOG/backend_outer/common"
)

func SeatRelease(w http.ResponseWriter, r *http.Request) {
	common.SetupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	seatSessionKey := r.Header.Get("Authorization")
	log.Println("the seat session key is", seatSessionKey)
	fmt.Println("Endpoint Hit: SeatRegister")
	releaseSeat(w, r, seatSessionKey)
}

func releaseSeat(w http.ResponseWriter, r *http.Request, sessionAuth string) {
	urlSeatRelease := "http://localhost:5001/powercubicle/v1/db/seat/release"

	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	fmt.Println("auth", sessionAuth)
	req, err := http.NewRequest(http.MethodPost, urlSeatRelease, nil)
	req.Header.Add("Authorization", sessionAuth)

	if err != nil {
		log.Fatal(err)
	}

	resp, getErr := cubeClient.Do(req)

	if resp.StatusCode == 400 || resp.StatusCode == 401 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		err_msg := string(body)
		log.Printf(err_msg)
		http.Error(w, err_msg, resp.StatusCode)
		return
	}

	if getErr != nil {
		log.Fatal(getErr)
	}
}
