package seat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	log.Println(seatCode)
	log.Println("Endpoint Hit: SeatRegister")
	registSeat(w, r, seatCode, seatSessionKey)
}

func registSeat(w http.ResponseWriter, r *http.Request, seatNumber string, sessionAuth string) {
	urlUserRegister := "http://localhost:5001/powercubicle/v1/db/seat/register"

	seatNumber = "02005"
	seatNumber = strings.TrimPrefix(seatNumber, "WS02.")
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

	res := map[string]string{"seat_number": seatNumber}
	json.NewEncoder(w).Encode(res)
}
