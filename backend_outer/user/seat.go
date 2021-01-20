package user

import (
	"encoding/json"
	"fmt"
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
	seatSessionKey := r.Header.Get("Authorization")
	fmt.Println("Ths session key is: ", seatSessionKey)
	userSeatInfo(w, r, seatSessionKey)
}

func userSeatInfo(w http.ResponseWriter, r *http.Request, sessionAuth string) {
	urlUserSeat := "http://222.186.160.104:5001/powercubicle/v1/db/user/seat"

	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, err := http.NewRequest(http.MethodGet, urlUserSeat, nil)
	req.Header.Add("Authorization", sessionAuth)

	if err != nil {
		log.Fatal(err)
	}

	res, getErr := cubeClient.Do(req)

	switch res.StatusCode {
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
	json.NewEncoder(w).Encode(userSeat)
}
