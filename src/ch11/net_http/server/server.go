package server

import (
	"encoding/json"
	"net/http"
	"net_http/middleware"
	"time"
)

// HelloHandler handles requests and sends a JSON response
type HelloHandler struct{}

// ServeHTTP implements the Handler interface for HelloHandler
func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Hello!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Server starts the HTTP server
func Server() {
	Handler := middleware.RequestTimer(HelloHandler{})

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		//Handler:      HelloHandler{},
		Handler: Handler,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
