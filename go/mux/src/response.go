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
	write(r.w, data)

	(*r.w).Header().Set("Content-Type", "application/json")
}

func (r *HTTPResponse) OkText(text string) {
	write(r.w, []byte(text))

	(*r.w).Header().Set("Content-Type", "text/plain")
}

func (r *HTTPResponse) CreatedJSON(data []byte) {
	write(r.w, data)

	(*r.w).WriteHeader(http.StatusCreated)
	(*r.w).Header().Set("Content-Type", "application/json")
}

func (r *HTTPResponse) BadRequestText(text string) {
	write(r.w, []byte(text))

	(*r.w).WriteHeader(http.StatusBadRequest)
	(*r.w).Header().Set("Content-Type", "text/plain")
}

func (r *HTTPResponse) Error(err error) {
	http.Error(*r.w, err.Error(), http.StatusInternalServerError)
}
