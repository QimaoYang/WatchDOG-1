package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ChrisLi03/WatchDOG/backend_outer/common"
)

type loginInfo struct {
	Name     string `json : "name"`
	Password string `json : "password"`
}

type loginSessionKey struct {
	Session_key string `json : "session_key"`
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	common.SetupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	body, readErr := ioutil.ReadAll(r.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	UserLoginInfo := loginInfo{}

	jsonErr := json.Unmarshal(body, &UserLoginInfo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// need further logic handler
	log.Printf("[watch dog] The user login info is %v", UserLoginInfo)
	getUserSessionKey(w, r, &UserLoginInfo)
}

func getUserSessionKey(w http.ResponseWriter, r *http.Request, userLoginInfo *loginInfo) {
	urlUserRegister := "http://localhost:5001/powercubicle/v1/db/user/login"

	userStat := map[string]string{
		"username": userLoginInfo.Name,
		"password": userLoginInfo.Password,
	}

	jsonString, jsonErr := json.Marshal(userStat)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("User info is %v", userStat)
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

	userSessionKey := loginSessionKey{}

	jsonErrSecond := json.Unmarshal(body, &userSessionKey)

	if jsonErrSecond != nil {
		log.Fatal(jsonErr)
	}

	// need further logic handler
	log.Printf("[watch dog] The user login session key is %v", userSessionKey)
	if userSessionKey.Session_key == "" {
		http.Error(w, "Bad request - wrong username or password", 400)
	} else {
		json.NewEncoder(w).Encode(userSessionKey)
	}
}
