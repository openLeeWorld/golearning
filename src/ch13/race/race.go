package race

import (
	"sync"
)

// 데이터 경쟁을 찾는다면 찾은 곳 주위에 적절한 잠금을 추가해야 한다.
func getCounter() int {
	var counter int
	var wg sync.WaitGroup
	wg.Add(5) // 고루틴이 Done을 통해 줄이는 카운터 5
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				counter++
				//time.Sleep(1000) // 1ms, 고루틴에서 sleep으로 해결 되진 않는다.
			}
			wg.Done() // 일 마치고 카운터 줄임
		}()
	}
	wg.Wait() // 카운터 0까지 메인 고루틴 대기
	return counter
}
