package fsq

import "fmt"

type ApiError struct {
	Code int
	Type string
	Msg  string
}

func (e *ApiError) Error() string {
	if e.Msg == "" {
		return fmt.Sprintf("[4SQ %d]: %s", e.Code, e.Msg)
	}
	return fmt.Sprintf("[4SQ %d]: %s - %s", e.Code, e.Type, e.Msg)
}

func (e *ApiError) IsAuthError() bool {
	return e.Code == 401
}
