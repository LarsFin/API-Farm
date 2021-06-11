package apifarm

import "net/http"

type Response interface {
	OkText(string)
}

type HTTPResponse struct {
	w *http.ResponseWriter
}

func NewHTTPResponse(w *http.ResponseWriter) *HTTPResponse {
	return &HTTPResponse{
		w,
	}
}

func (r *HTTPResponse) OkText(text string) {
	_, err := (*r.w).Write([]byte(text))

	if err != nil {
		http.Error(*r.w, err.Error(), http.StatusInternalServerError)
	}

	(*r.w).Header().Set("Content-Type", "text/plain")
}
