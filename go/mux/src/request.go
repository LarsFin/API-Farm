package apifarm

import (
	"io/ioutil"
	"net/http"
)

type Request interface {
	GetBody() ([]byte, error)
}

type HTTPRequest struct {
	r *http.Request
}

func NewHTTPRequest(r *http.Request) *HTTPRequest {
	return &HTTPRequest{
		r,
	}
}

func (r *HTTPRequest) GetBody() ([]byte, error) {
	return ioutil.ReadAll(r.r.Body)
}
