package server

import (
	"encoding/json"
	"net/http"

	"whitetrefoil.com/checkin/fsq"
	"whitetrefoil.com/checkin/server/jr"
)

type checkLoginReqJson struct {
	Code string `json:"code"`
}

type checkLoginResJson struct {
	Token string                     `json:"token"`
	User  *fsq.GetUserDetailResponse `json:"user"`
}

func checkLogin(w http.ResponseWriter, r *http.Request) {
	req := &checkLoginReqJson{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		jr.Json400(w, err.Error())
		return
	}
	if req.Code == "" {
		jr.Json400(w, "missing code")
		return
	}
	cfg := r.Context().Value("cfg").(*Config)

	tokenRes, err := fsq.AccessToken(cfg.AppId, cfg.AppSecret, cfg.Redirect, req.Code)
	if err != nil {
		jr.Json400(w, err.Error())
		return
	}

	userRes, err := fsq.GetUserDetail(tokenRes.AccessToken)
	if err != nil {
		if err, ok := err.(*fsq.ApiError); ok {
			jr.Json(w, err.Code, err.Error())
			return
		}
		jr.Json400(w, err.Error())
		return
	}

	jr.Json200(w, &checkLoginResJson{
		Token: tokenRes.AccessToken,
		User:  userRes,
	})
}
