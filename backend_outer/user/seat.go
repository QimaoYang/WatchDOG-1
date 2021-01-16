package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ChrisLi03/WatchDOG/backend_outer/common"
)

type Seat struct {
	Seat string `json : "seat"`
}

func UserSeat(w http.ResponseWriter, r *http.Request) {
	common.SetupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	seatSessionKey := r.Header.Get("Authorization")
	log.Println("the seat session key is", seatSessionKey)
	userSeatInfo(w, r, seatSessionKey)
}

func userSeatInfo(w http.ResponseWriter, r *http.Request, sessionAuth string) {
	urlUserSeat := "http://localhost:5001/powercubicle/v1/db/user/seat"

	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, err := http.NewRequest(http.MethodGet, urlUserSeat, nil)
	req.Header.Add("Authorization", sessionAuth)

	if err != nil {
		log.Fatal(err)
	}

	res, getErr := cubeClient.Do(req)

	if res.StatusCode == 400 || res.StatusCode == 401 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		err_msg := string(body)
		log.Printf(err_msg)
		http.Error(w, err_msg, res.StatusCode)
		return
	}

	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	userSeat := Seat{}

	jsonErrSecond := json.Unmarshal(body, &userSeat)

	if jsonErrSecond != nil {
		log.Fatal(jsonErrSecond)
	}

	// need further logic handler
	log.Printf("[watch dog] The user's reserved seat is %v", userSeat)
	// if userSessionKey.Session_key == "" {
	// 	http.Error(w, "Bad request - user already exists!", 400)
	// } else {
	// 	json.NewEncoder(w).Encode(userSessionKey)
	// }
	json.NewEncoder(w).Encode(userSeat)
}
