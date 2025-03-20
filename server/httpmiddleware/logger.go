package httpmiddleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger(level zerolog.Level) func(http.Handler) http.Handler {
	logger := log.Logger
	f := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			event := logger.WithLevel(level)
			var rw *respWrapper

			if event.Enabled() {
				event = event.Str("http.request.remote_addr", r.RemoteAddr)
				event = event.Str("http.request.method", r.Method)
				event = event.Str("http.request.url", r.URL.Path)
				event = event.Str("http.request.host", r.Host)
				event = event.Str("http.request.proto", r.Proto)

				for header, values := range r.Header {
					event = event.Strs(fmt.Sprintf("http.request.headers.%s", header), values)
				}

				for k, v := range r.URL.Query() {
					event = event.Strs(fmt.Sprintf("http.request.query.%s", k), v)
				}

				body, err := io.ReadAll(r.Body)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
				event = event.Bytes("http.request.body", body)
				r.Body = io.NopCloser(bytes.NewBuffer(body))

				rw = newRespWrapper(w)
				defer rw.Flush()
				w = rw
			}

			start := time.Now()
			next.ServeHTTP(w, r)
			finish := time.Now()

			if event.Enabled() {
				event = event.Int("http.response.status_code", rw.httpStatusCode)
				event = event.Str("http.response.body", rw.response.String())
				dur := finish.Sub(start)
				event = event.Dur("http.duration", dur)

				event.Msg(fmt.Sprintf("%s %s - %d %s", r.Method, r.URL.Path, rw.httpStatusCode, dur))
			}
		}
		return http.HandlerFunc(fn)
	}
	return f
}
