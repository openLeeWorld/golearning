package main

import (
	"fmt"
	"time"
)

func main() { // main 고루틴
	t, err := time.Parse("2006-01-02 15:04:05 -0700", "2016-03-13 00:00:00 +0000")
	// time.Parse(포맷, 실제 변환 문자열)로 문자열을 time.Time 타입으로 변환
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
	// time.Format("time.Time 포맷")으로 time.Time에서 문자열로 변환
}
