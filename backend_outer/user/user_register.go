package user

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type userInfo struct {
	Name     string `json : "name"`
	Password string `json : "password"`
}

type sessionKey struct {
	Session_key string `json : "session_key"`
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, readErr := ioutil.ReadAll(r.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	UserRegisterInfo := userInfo{}

	jsonErr := json.Unmarshal(body, &UserRegisterInfo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// need further logic handler
	log.Printf("[watch dog] The user register info is %v", UserRegisterInfo)
	getSessionKey(w, r, &UserRegisterInfo)
}

func getSessionKey(w http.ResponseWriter, r *http.Request, userRegisterInfo *userInfo) {
	urlUserRegister := "http://localhost:5001/powercubicle/v1/db/user/register"

	userStat := map[string]string{
		"username": userRegisterInfo.Name,
		"password": userRegisterInfo.Password,
	}

	jsonString, jsonErr := json.Marshal(userStat)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	log.Printf("User info is %v", userStat)
	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, err := http.NewRequest(http.MethodPost, urlUserRegister, bytes.NewBuffer(jsonString))

	if err != nil {
		log.Fatal(err)
	}

	res, getErr := cubeClient.Do(req)

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

	userSessionKey := sessionKey{}

	jsonErrSecond := json.Unmarshal(body, &userSessionKey)

	if jsonErrSecond != nil {
		log.Fatal(jsonErr)
	}

	// need further logic handler
	log.Printf("[watch dog] The user register session key is %v", userSessionKey)
	if userSessionKey.Session_key == "" {
		http.Error(w, "Bad request - user already exists!", 400)
	} else {
		json.NewEncoder(w).Encode(userSessionKey)
	}
}
