package main

import (
	"fmt" // 내장 모듈
)

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	// done 채널은 닫는 용도
	cancel := func() {
		close(done)
	} // done채널을 닫는 클로저 반환용
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done: // 읽을 것이 더이상 없을 때, 갑자기 채널 읽기가 중단될 때
				return
			default: // 기본으로 개수만큼 씀
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}

func main() { // main 고루틴
	ch, cancel := countTo(10)
	for i := range ch {
		if i > 5 {
			break
		} // 채널에 있는 것을 다 못 읽으면 채널은 정지됨
		fmt.Println(i)
	}
	cancel() // 수동으로 채널 종료하기
}
