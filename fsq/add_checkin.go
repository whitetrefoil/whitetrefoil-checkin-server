package fsq

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type scoreDetail struct {
	Icon    string `json:"icon"`
	Message string `json:"message"`
	Points  int    `json:"points"`
}

type addCheckinRes struct {
	Meta     *meta `json:"meta"`
	Response *struct {
		Checkin *struct {
			CheckinShortUrl string `json:"checkinShortUrl"`
			IsMayor         bool   `json:"isMayor"`
			Score           *struct {
				Total  int            `json:"total"`
				Scores []*scoreDetail `json:"scores"`
			} `json:"score"`
		} `json:"checkin"`
	} `json:"response"`
}

type AddCheckinResponse struct {
	IsMayor bool           `json:"is_mayor"`
	Score   int            `json:"score"`
	Reasons []*scoreDetail `json:"reasons"`
	Url     string         `json:"url"`
}

func AddCheckin(token string, venueId string, shout string, lat float64, lon float64, acc float64, alt float64) (*AddCheckinResponse, error) {
	u, _ := url.Parse("https://api.foursquare.com/v2/checkins/add")
	q := u.Query()
	q.Set("v", ApiVersion)
	q.Set("oauth_token", token)
	q.Set("venueId", venueId)
	q.Set("ll", fmt.Sprintf("%f,%f", lat, lon))
	q.Set("llAcc", fmt.Sprintf("%f", acc))
	q.Set("alt", fmt.Sprintf("%f", alt))
	if shout != "" {
		q.Set("shout", shout)
	}
	u.RawQuery = q.Encode()

	log.Printf("POST %s", u)
	resp, err := http.Post(u.String(), "", nil)
	if err != nil {
		return nil, err
	}

	res := &addCheckinRes{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, &ApiError{resp.StatusCode, res.Meta.ErrorType, res.Meta.ErrorDetail}
	}

	details := make([]*scoreDetail, len(res.Response.Checkin.Score.Scores))
	for i, detail := range res.Response.Checkin.Score.Scores {
		details[i] = detail
	}

	return &AddCheckinResponse{
		IsMayor: res.Response.Checkin.IsMayor,
		Score:   res.Response.Checkin.Score.Total,
		Reasons: details,
		Url:     res.Response.Checkin.CheckinShortUrl,
	}, nil
}
