package apifarm

import "net/http"

type response interface {
	OkText(string)
}

type httpResponse struct {
	w *http.ResponseWriter
}

func NewHttpResponse(w *http.ResponseWriter) *httpResponse {
	return &httpResponse{
		w,
	}
}

func (r *httpResponse) OkText(text string) {
	_, err := (*r.w).Write([]byte(text))

	if err != nil {
		http.Error(*r.w, err.Error(), http.StatusInternalServerError)
	}

	(*r.w).Header().Set("Content-Type", "text/plain")
}
