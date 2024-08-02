package main

import (
	"fmt" // 내장 모듈
)

func main() { // main 고루틴
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2 // 여기서 block됨
		fmt.Println("Go routine", v, v2)
	}() // gorouting 선언

	v := 2
	var v2 int
	select { // select문은 가용한 채널을 랜덤으로 선택해서 데드락 회피
	case ch2 <- v: // ch2에 write
		fmt.Println("Send to ch2")
	case v2 = <-ch1: //ch1에서 read
		fmt.Println("Received from ch1")
	}
	fmt.Println("Main routine", v, v2)
}

/* // 아래는 데드락: all goroutinesa are asleep - deadlock!
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v // 고루틴은 ch1이 읽힐 때까지 진행 x
		v2 := <-ch2
		fmt.Println(v, v2)
	}() // gorouting 선언
	v := 2
	var v2 int
	ch2 <- v: // ch2에 write, ch2이 읽힐 때까지 진행 x
	v2 = <-ch1: //ch1에서 read
	fmt.Println(v, v2)
}
*/

/* // 아래는 for-select 루프로 처리하는 경우이다. (계속 채널 수신,송신?)
for {
	select {
	case <- done: // done 채널
		return
	case v := <- ch:
		fmt.Println(v)
	//default: // 매 루프에서 모든 케이스가 다 작동 못할 때 시도
	// 일반 select에서는 default로 무조건 선택되므로 주의
	}
}
*/
