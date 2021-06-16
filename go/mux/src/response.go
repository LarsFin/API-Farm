package apifarm

import "net/http"

type Response interface {
	OkJSON([]byte)
	OkText(string)
	CreatedJSON([]byte)
	Error(error)
}

type HTTPResponse struct {
	w *http.ResponseWriter
}

func NewHTTPResponse(w *http.ResponseWriter) *HTTPResponse {
	return &HTTPResponse{
		w,
	}
}

func (r *HTTPResponse) OkJSON(data []byte) {
	_, err := (*r.w).Write(data)

	if err != nil {
		http.Error(*r.w, err.Error(), http.StatusInternalServerError)
	}

	(*r.w).Header().Set("Content-Type", "application/json")
}

func (r *HTTPResponse) OkText(text string) {
	_, err := (*r.w).Write([]byte(text))

	if err != nil {
		http.Error(*r.w, err.Error(), http.StatusInternalServerError)
	}

	(*r.w).Header().Set("Content-Type", "text/plain")
}

func (r *HTTPResponse) CreatedJSON(data []byte) {
	_, err := (*r.w).Write(data)

	if err != nil {
		http.Error(*r.w, err.Error(), http.StatusInternalServerError)
	}

	(*r.w).WriteHeader(http.StatusCreated)
	(*r.w).Header().Set("Content-Type", "application/json")
}

func (r *HTTPResponse) Error(err error) {
	http.Error(*r.w, err.Error(), http.StatusInternalServerError)
}
