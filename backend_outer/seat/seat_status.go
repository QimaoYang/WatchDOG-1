package seat

import (
	"fmt"
	"net/http"
)

// type SeatInfo struct {
// 	Status string                 `json: status`
// 	Data   map[string]interface{} `json: data`
// }

func RetrieveAllSeatStatus(w http.ResponseWriter, r *http.Request) {
	// var seats SeatInfo
	// seats.Status = "success"
	// seats.Data = `{
	// 	"region": "all"
	// 	"seats": [
	// 		{"A1" : "availble"},
	// 		{"B2": "Zhang3"}
	// 	]
	// }`
	seatInfo := map[string]interface{}{
		"status": "success",
		"data": struct {
			name []map[string]string
		}{name: [{"id": "A", "rest": 10}]},
	}
	
	fmt.Println(seatInfo)
	// responseData := []byte(seatJson)
	// fmt.Println(responseData)
	// json.NewEncoder(w).Encode(seatInfo)
	// fmt.Println("Endpoint Hit: RetrieveAllSeatStatus")
}
