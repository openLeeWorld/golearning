package middleware

import (
	"log"
	"net/http"
	"time"
)

// RequestTimer middleware for time elasped checking
func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s:%v", r.URL.Path, end.Sub(start))
	})
}
