package apifarm

import "net/http"

type Response interface {
	OkJSON([]byte)
	OkText(string)
	CreatedJSON([]byte)

	BadRequestText(string)

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

func write(w *http.ResponseWriter, data []byte) {
	_, err := (*w).Write(data)

	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
	}
}

func (r *HTTPResponse) OkJSON(data []byte) {
	(*r.w).Header().Set("Content-Type", "application/json")

	write(r.w, data)
}

func (r *HTTPResponse) OkText(text string) {
	(*r.w).Header().Set("Content-Type", "text/plain")

	write(r.w, []byte(text))
}

func (r *HTTPResponse) CreatedJSON(data []byte) {
	(*r.w).Header().Set("Content-Type", "application/json")
	(*r.w).WriteHeader(http.StatusCreated)

	write(r.w, data)
}

func (r *HTTPResponse) BadRequestText(text string) {
	(*r.w).Header().Set("Content-Type", "text/plain")
	(*r.w).WriteHeader(http.StatusBadRequest)

	write(r.w, []byte(text))
}

func (r *HTTPResponse) Error(err error) {
	http.Error(*r.w, err.Error(), http.StatusInternalServerError)
}
