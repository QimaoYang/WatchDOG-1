package seat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ChrisLi03/WatchDOG/backend_outer/management"
)

type QrCode struct {
	Encryted string `json:"encrypted_qrcode"`
}

func SeatRegister(w http.ResponseWriter, r *http.Request) {
	seatSessionKey := r.Header.Get("Auth")
	log.Println("the seat session key is", seatSessionKey)
	var p QrCode

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	seatCode := management.DecryptCode(p.Encryted)
	res := map[string]string{"seat_number": seatCode}
	json.NewEncoder(w).Encode(res)
	fmt.Println(seatCode)
	fmt.Println("Endpoint Hit: SeatRegister")
	registSeat(w, r, seatCode, seatSessionKey)
}

func registSeat(w http.ResponseWriter, r *http.Request, seatNumber string, sessionAuth string) {
	urlUserRegister := "http://localhost:5001/powercubicle/v1/db/seat/register"

	seatNumber = "02005"
	seatCode := map[string]string{
		"seat_code": seatNumber,
	}

	jsonString, jsonErr := json.Marshal(seatCode)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	log.Printf("User info is %v", seatCode)
	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	fmt.Println("auth", sessionAuth)
	req, err := http.NewRequest(http.MethodPost, urlUserRegister, bytes.NewBuffer(jsonString))
	req.Header.Add("Authorization", sessionAuth)

	if err != nil {
		log.Fatal(err)
	}

	_, getErr := cubeClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}
}
