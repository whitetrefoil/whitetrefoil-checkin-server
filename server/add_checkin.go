package server

import (
	"encoding/json"
	"net/http"

	"whitetrefoil.com/checkin/fsq"
	"whitetrefoil.com/checkin/server/jr"
)

type addCheckinReq struct {
	VenueId   string  `json:"venue_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Accuracy  float64 `json:"accuracy"`
	Altitude  float64 `json:"altitude"`
	Shout     string  `json:"shout"`
}

func addCheckin(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(string)
	req := &addCheckinReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		jr.Json400(w, err.Error())
		return
	}

	if req.VenueId == "" {
		jr.Json400(w, "\"venueId\" is required")
		return
	}

	result, err := fsq.AddCheckin(token, req.VenueId, req.Shout, req.Latitude, req.Longitude, req.Accuracy, req.Altitude)
	if err != nil {
		jr.Json400(w, err.Error())
		return
	}

	jr.Json201(w, result)
}
