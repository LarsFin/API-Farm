package apifarm

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Request interface {
	GetBody() ([]byte, error)
	GetParam(string) string
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

func (r *HTTPRequest) GetParam(p string) string {
	params := mux.Vars(r.r)
	return params[p]
}
