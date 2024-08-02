package servers

import (
	"net/http"
	"net/http/httptest" // 원격 서버와 통신하는 코드에 대한 유닛 테스트
	"time"
)

// SlowServer sleeps 2 seconds and response "slow response"
func SlowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Slow response"))
	}))
	return s
}

// FastServer checks the query parameter "error" and if it is "true", returns "error"
func FastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		if r.URL.Query().Get("error") == "true" {
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	}))
	return s
}
