package seat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChrisLi03/WatchDOG/backend_outer/management"
)

type QrCode struct {
	Encryted string `json:"encrypted_qrcode"`
}

func SeatRegister(w http.ResponseWriter, r *http.Request) {
	var p QrCode

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(p.Encryted)
	seatCode := management.DecryptCode(p.Encryted)
	res := map[string]string{"seat_number": seatCode}
	json.NewEncoder(w).Encode(res)
	fmt.Println(seatCode)
	fmt.Println("Endpoint Hit: SeatRegister")
}
