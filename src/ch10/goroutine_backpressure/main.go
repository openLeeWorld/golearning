package main

import (
	"goroutine_backpressure/backpressure"
	"net/http"
	"time"
)

func doThingThatShouldBeLimited() string {
	time.Sleep(3 * time.Second)
	return "done"
}

func main() { // main 고루틴
	pg := backpressure.New(10) // 고루틴 토큰 limit이 10개인 채널 생성, 해당 타입의 포인터
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
