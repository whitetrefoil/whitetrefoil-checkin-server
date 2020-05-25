package server

import (
	"fmt"
	"net/http"
	"strconv"

	"whitetrefoil.com/checkin/fsq"
	"whitetrefoil.com/checkin/server/jr"
)

func searchVenues(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(string)

	if err := r.ParseForm(); err != nil {
		jr.Json400(w, err)
		return
	}
	name := r.Form.Get("name")
	latitudeStr := r.Form.Get("latitude")
	longitudeStr := r.Form.Get("longitude")
	accuracyStr := r.Form.Get("accuracy")
	altitudeStr := r.Form.Get("altitude")
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	accuracy, err := strconv.ParseFloat(accuracyStr, 64)
	altitude, err := strconv.ParseFloat(altitudeStr, 64)

	if err != nil {
		jr.Json400(w, fmt.Sprintf("missing required param: %s", err))
		return
	}

	result, err := fsq.SearchVenues(token, latitude, longitude, accuracy, altitude, name)
	if err != nil {
		if err, ok := err.(*fsq.ApiError); ok {
			if err.IsAuthError() {
				jr.Json401(w, err.Error())
				return
			}
			jr.Json(w, err.Code, err.Error())
			return
		}
		jr.Json400(w, err.Error())
		return
	}

	jr.Json200(w, result)
}
