package main

import (
	"fmt" // 내장 모듈
)

func main() { // main 고루틴
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		//v := v // 일부로 shadowing
		go func(val int) { //고루틴이 동시에 실행되면 언제 채널에 읽고 쓸 지 알 수 없음
			ch <- val * 2
		}(v)
	} // 고루틴 동기화나 실행 순서 조정 등으로 쓰는 순서를 맞출 수도 있으나
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
	close(ch)
}
