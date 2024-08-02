package testmain

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

// TestMain 함수가 있는 패키지에서 go test를 실행하면 테스트 함수 호출
func TestMain(m *testing.M) {
	fmt.Println("Set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal) // 0은 모든 테스트가 통과했음
} // TestMain은 패키지 별로 하나만 가짐
// 1. DB같은 외부 저장소에 있는 데이터 설정이 필요한 경우
// 2. 초기화 될 필요가 있는 패키지 레벨 변수에 의존적인 코드가 테스트 될 경우

func TestFirst(t *testing.T) {
	fmt.Println("Test first uses stuff set up in TestMain", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("Test Second also uses stuff set up in TestMain", testTime)
}
