package server

import (
	"net/http"
	"net/url"

	"whitetrefoil.com/checkin/server/jr"
)

type getLoginJson struct {
	Url string `json:"url"`
}

func getLoginUrl(w http.ResponseWriter, r *http.Request) {
	cfg := r.Context().Value("cfg").(*Config)
	u, _ := url.Parse("https://foursquare.com/oauth2/authenticate")
	q := u.Query()
	q.Set("client_id", cfg.AppId)
	q.Set("response_type", "code")
	q.Set("redirect_uri", cfg.Redirect)
	u.RawQuery = q.Encode()

	jr.Json200(w, &getLoginJson{u.String()})
}
