package backpressure

import "errors"

type PressureGauge struct {
	ch chan struct{}
} // 배압이 있는 채널 타입 선언

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit) // capacity가 있는 채널
	for i := 0; i < limit; i++ {
		ch <- struct{}{} // ch에 빈 구조체 리터럴(필드 없고 크기 0) write
	} // 토큰의 개수를 10개로 제한
	return &PressureGauge{ // PressureGauge의 구조체 포인터를 반환
		ch: ch,
	}
}

// 모든 고루틴은 Process함수를 사용하길 원한다. f를 실행하기 위해
func (pg *PressureGauge) Process(f func()) error { // pg의 메서드
	select { // 채널로부터 토큰을 읽음
	case <-pg.ch: // ch에서 읽어올 게 있으면 실행
		f()                 // 함수 실행
		pg.ch <- struct{}{} // 채널에 다시 토큰 등록
		return nil          // 이상 없이 함수 실행
	default: // 토큰을 읽을 수 없다면, 오류 반환
		return errors.New("no more capacity")
	}
}
