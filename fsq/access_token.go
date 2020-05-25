package fsq

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type ResJson struct {
	AccessToken string `json:"access_token"`
	Error       string `json:"error"`
}

func AccessToken(id string, secret string, redirect string, code string) (*ResJson, error) {
	u, _ := url.Parse("https://foursquare.com/oauth2/access_token")
	q := u.Query()
	q.Set("client_id", id)
	q.Set("client_secret", secret)
	q.Set("grant_type", "authorization_code")
	q.Set("redirect_uri", redirect)
	q.Set("code", code)
	u.RawQuery = q.Encode()

	log.Printf("GET %s", u)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	res := &ResJson{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, &ApiError{resp.StatusCode, res.Error, ""}
	}

	return res, nil
}
