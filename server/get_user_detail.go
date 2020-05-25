package server

import (
	"net/http"

	"whitetrefoil.com/checkin/fsq"
	"whitetrefoil.com/checkin/server/jr"
)

func getUserDetail(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(string)

	res, err := fsq.GetUserDetail(token)
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

	jr.Json200(w, res)
}
