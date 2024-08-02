package main

import (
	"net_http/client"
	"net_http/server"
	"time"
)

func main() { // main 고루틴
	go server.Server() // 서버 켜놓기

	time.Sleep(1 * time.Second) // 서버가 시작되길 기다림

	for i := 0; i < 10; i++ {
		client.Client()             // 클라이언트 요청 실행
		time.Sleep(1 * time.Second) // 서버가 시작되길 기다림
	}
}
