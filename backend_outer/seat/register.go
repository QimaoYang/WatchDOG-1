package seat

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ChrisLi03/WatchDOG/backend_outer/common"
	"github.com/ChrisLi03/WatchDOG/backend_outer/management"
)

type QrCode struct {
	Encryted string `json:"encrypted_qrcode"`
}

func SeatRegister(w http.ResponseWriter, r *http.Request) {
	common.SetupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	seatSessionKey := r.Header.Get("Authorization")
	log.Println("the seat session key is", seatSessionKey)
	var p QrCode

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	seatCode := management.DecryptCode(p.Encryted)
	// res := map[string]string{"seat_number": seatCode}
	// json.NewEncoder(w).Encode(res)
	log.Println("[WD] Raw seat code is", seatCode)
	registSeat(w, r, seatCode, seatSessionKey)
}

func registSeat(w http.ResponseWriter, r *http.Request, seatNumber string, sessionAuth string) {
	urlUserRegister := "http://192.168.242.158:5001/powercubicle/v1/db/seat/register"

	seatNumber = strings.TrimPrefix(seatNumber, "WS02.")
	log.Println("[WD] Start booking seat ", seatNumber)

	seatCode := map[string]string{
		"seat_code": seatNumber,
	}

	jsonString, jsonErr := json.Marshal(seatCode)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, err := http.NewRequest(http.MethodPost, urlUserRegister, bytes.NewBuffer(jsonString))
	req.Header.Add("Authorization", sessionAuth)

	if err != nil {
		log.Fatal(err)
	}

	resp, getErr := cubeClient.Do(req)

	switch resp.StatusCode {
	case 400:
		errMsg := strings.Join([]string{"The seat", seatNumber, "has been booked"}, " ")
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

	res := map[string]string{"seat_number": seatNumber}
	json.NewEncoder(w).Encode(res)
}
