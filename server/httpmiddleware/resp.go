package httpmiddleware

import (
	"net/http"
	"strings"
)

type respWrapper struct {
	rw             http.ResponseWriter
	httpStatusCode int
	response       strings.Builder
}

func newRespWrapper(rw http.ResponseWriter) *respWrapper {
	return &respWrapper{rw: rw, response: strings.Builder{}}
}

func (r *respWrapper) Header() http.Header {
	return r.rw.Header()
}

func (r *respWrapper) Write(bytes []byte) (int, error) {
	r.response.Write(bytes)
	return r.rw.Write(bytes)
}

func (r *respWrapper) WriteHeader(statusCode int) {
	r.httpStatusCode = statusCode
	r.rw.WriteHeader(statusCode)
}

func (r *respWrapper) Flush() {
	r.response.Reset()
}
