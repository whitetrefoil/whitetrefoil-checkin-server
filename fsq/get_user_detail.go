package fsq

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type getUserDetailRes struct {
	Meta     *meta `json:"meta"`
	Response *struct {
		User *struct {
			ID        string `json:"id"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Photo     *struct {
				// prefix + size + suffix
				// size: 36x36, 100x100, 300x300, and 500x500
				Prefix string `json:"prefix"`
				Suffix string `json:"suffix"`
			} `json:"photo"`
		} `json:"user"`
	} `json:"response"`
}

type GetUserDetailResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Photo     *struct {
		// prefix + size + suffix
		// size: 36x36, 100x100, 300x300, and 500x500
		Prefix string `json:"prefix"`
		Suffix string `json:"suffix"`
	} `json:"photo"`
}

func GetUserDetail(token string) (*GetUserDetailResponse, error) {
	u, _ := url.Parse("https://api.foursquare.com/v2/users/self")
	q := u.Query()
	q.Set("v", API_VERSION)
	q.Set("oauth_token", token)
	u.RawQuery = q.Encode()

	log.Printf("GET %s", u)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	res := &getUserDetailRes{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("[4SQ %d]: %s - %s", resp.StatusCode, res.Meta.ErrorType, res.Meta.ErrorDetail)
	}

	return &GetUserDetailResponse{
		ID:        res.Response.User.ID,
		FirstName: res.Response.User.FirstName,
		LastName:  res.Response.User.LastName,
		Photo: &struct {
			Prefix string `json:"prefix"`
			Suffix string `json:"suffix"`
		}{
			Prefix: res.Response.User.Photo.Prefix,
			Suffix: res.Response.User.Photo.Suffix,
		},
	}, nil
}
