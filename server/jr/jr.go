package jr

import (
	"encoding/json"
	"log"
	"net/http"
)

type BaseResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

func setHeader(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func jsonErr(w http.ResponseWriter, err error) {
	if err != nil {
		log.Printf("[ERR] json response error: %s\n", err)
	}
	setHeader(w, 500)
	w.Write([]byte("{\"code\":500}"))
}

func Json(w http.ResponseWriter, code int, v interface{}) {
	body := &BaseResponse{
		code,
		v,
	}
	text, err := json.Marshal(body)
	if err != nil {
		jsonErr(w, err)
		return
	}

	setHeader(w, code)
	w.Write(text)
}

func Json200(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusOK, v)
}

func Json201(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusCreated, v)
}

func Json400(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusBadRequest, v)
}

func Json401(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusUnauthorized, v)
}

func Json404(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusNotFound, v)
}

func Json405(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusMethodNotAllowed, v)
}

func Json409(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusConflict, v)
}

func Json500(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusInternalServerError, v)
}

func Json502(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusBadGateway, v)
}

func Json504(w http.ResponseWriter, v interface{}) {
	Json(w, http.StatusGatewayTimeout, v)
}
