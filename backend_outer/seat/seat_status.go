package seat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type seatStatus struct {
	Seats []map[string]int `json:"seats"`
}

func RetrieveAllSeatStatus(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query()
	if reg, ok := v1["region"]; ok {
		retrieveRegionSeatStatus(w, r, reg)
	} else {
		s := []string{"1", "2", "3"}
		retrieveRegionSeatStatus(w, r, s)
	}
}

func retrieveRegionSeatStatus(w http.ResponseWriter, r *http.Request, selectedRegion []string) {
	seatInfo := []map[string]int{}
	for _, v := range selectedRegion {
		fmt.Println(v)
		response, err := http.Get(strings.Join([]string{"http://localhost:5001/Reservation/available", v, "count"}, "/"))

		if err != nil {
			fmt.Print(err.Error())
		}

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		var responseObject map[string]int
		errJson := json.Unmarshal(responseData, &responseObject)
		if errJson != nil {
			fmt.Println(errJson)
		}

		fmt.Println(string(responseData))
		fmt.Println(responseObject)
		seatInfo = append(seatInfo, map[string]int{v: responseObject["available"]})
	}
	fmt.Println(seatInfo)
	m := seatStatus{Seats: seatInfo}
	fmt.Println(m)
	json.NewEncoder(w).Encode(m)
}
