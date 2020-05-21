package fsq

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type searchVenuesRes struct {
	Meta     *meta `json:"meta"`
	Response *struct {
		Venues []*struct {
			Id       string `json:"id"`
			Name     string `json:"name"`
			Location *struct {
				Address  string  `json:"address"`
				Distance float64 `json:"distance"`
			} `json:"location"`
		} `json:"venues"`
	} `json:"response"`
}

type Venue struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Address  string  `json:"address"`
	Distance float64 `json:"distance"`
}

type SearchVenuesResponse = []*Venue

func SearchVenues(token string, lat float64, lon float64, acc float64, alt float64, name string) (*SearchVenuesResponse, error) {
	u, _ := url.Parse("https://api.foursquare.com/v2/venues/search")
	q := u.Query()
	q.Set("oauth_token", token)
	q.Set("v", API_VERSION)
	q.Set("intent", "checkin")
	q.Set("ll", fmt.Sprintf("%f,%f", lat, lon))
	q.Set("llAcc", fmt.Sprintf("%f", acc))
	q.Set("alt", fmt.Sprintf("%f", alt))
	if name != "" {
		q.Set("query", name)
	}
	q.Set("limit", "50")
	u.RawQuery = q.Encode()

	log.Printf("GET %s", u)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	res := &searchVenuesRes{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("[4SQ %d]: %s - %s", resp.StatusCode, res.Meta.ErrorType, res.Meta.ErrorDetail)
	}

	venues := make([]*Venue, len(res.Response.Venues))
	for i, v := range res.Response.Venues {
		venues[i] = &Venue{
			Id:       v.Id,
			Name:     v.Name,
			Address:  v.Location.Address,
			Distance: v.Location.Distance,
		}
	}

	return &venues, nil
}
