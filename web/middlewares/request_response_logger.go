package middlewares

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func RequestResponseLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Host
		// HTTP Method
		// Endpoint
		// From IP
		// Status Code
		// How many bytes the response
		// Latency
		start := time.Now()

		h.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Info().
			Str("host", getHost(r)).
			Str("method", getHTTPMethod(r)).
			Str("endpoint", getHTTPEndpoint(r)).
			Str("source", getRequesterIP(r)).
			Int("status_code", getHTTPStatusCode()).
			Str("size", fmt.Sprintf("%v bytes", responseBodySize(w, r))).
			Dur("latency_ms", duration).
			Send()
	})
}

func getHost(r *http.Request) string {
	if addr, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
		return addr.String()
	}

	return ""
}

func ipFromHostPort(hp string) string {
	h, _, err := net.SplitHostPort(hp)
	if err != nil {
		return ""
	}
	if len(h) > 0 && h[0] == '[' {
		return h[1 : len(h)-1]
	}
	return h
}

func getHTTPMethod(r *http.Request) string {
	return r.Method
}

func getHTTPEndpoint(r *http.Request) string {
	return r.URL.String()
}

func getRequesterIP(r *http.Request) string {
	return ipFromHostPort(r.RemoteAddr)
}

func getHTTPStatusCode() int {
	return http.StatusOK
}

func responseBodySize(w http.ResponseWriter, r *http.Request) int64 {
	return 0
}
