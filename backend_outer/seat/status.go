package seat

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type seatStatus struct {
	Data struct {
		Seat []map[string]int `json : "seat"`
	} `json : "data"`
}

type regionSeats struct {
	Data struct {
		Seat []map[string]string `json : "seat"`
	} `json : "data"`
}

func RetrieveAllSeatStatus(w http.ResponseWriter, r *http.Request) {

	v1 := r.URL.Query()
	if reg, ok := v1["region"]; ok {
		retrieveRegionStatus(w, r, reg[0])
	} else {
		retrieveOverallStatus(w, r)
	}
}

func retrieveOverallStatus(w http.ResponseWriter, r *http.Request) {
	urlSeats := "http://139.198.15.216:5001/powercubicle/v1/db/seat"

	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, urlSeats, nil)

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

	seatInfo := seatStatus{}

	jsonErr := json.Unmarshal(body, &seatInfo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// need further logic handler
	log.Printf("[watch dog] The overall seats info is %v", seatInfo)
	json.NewEncoder(w).Encode(seatInfo)
}

func retrieveRegionStatus(w http.ResponseWriter, r *http.Request, region string) {
	urlSeats := "http://139.198.15.216:5001/powercubicle/v1/db/seat"
	urlRegionSeats := strings.Join([]string{urlSeats, "?region=", region}, "")
	cubeClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, urlRegionSeats, nil)

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

	seatInfo := regionSeats{}

	jsonErr := json.Unmarshal(body, &seatInfo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	regionSeat := map[string]map[string]string{}
	seatDetails := map[string]string{}
	seatKeys := []string{}
	var seatNum string

	for _, reg := range seatInfo.Data.Seat {
		seatKeys = getKeys(reg)
		for _, key := range seatKeys {
			if key != "team" {
				seatNum = key
				break
			}
		}
		seatDetails["team"] = reg["team"]
		seatDetails["state"] = reg[seatNum]
		regionSeat[seatNum] = seatDetails
	}

	// need further logic handler
	log.Printf("[watch dog] The region %s seats info is %v", region, regionSeat)

	json.NewEncoder(w).Encode(regionSeat)
}

func getKeys(m map[string]string) []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}
