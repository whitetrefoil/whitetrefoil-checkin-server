package fsq

type meta struct {
	Code        int    `json:"code"`
	ErrorType   string `json:"errorType"`
	ErrorDetail string `json:"errorDetail"`
	RequestId   string `json:"requestId"`
}
