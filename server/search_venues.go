package server

import (
	"encoding/json"
	"net/http"

	"whitetrefoil.com/checkin/fsq"
	"whitetrefoil.com/checkin/server/jr"
)

type searchVenuesReq struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Accuracy  float64 `json:"accuracy"`
	Altitude  float64 `json:"altitude"`
}

func searchVenues(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(string)

	req := &searchVenuesReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		jr.Json400(w, err.Error())
		return
	}

	result, err := fsq.SearchVenues(token, req.Latitude, req.Longitude, req.Accuracy, req.Altitude, req.Name)
	if err != nil {
		jr.Json400(w, err.Error())
		return
	}

	jr.Json200(w, result)
}
