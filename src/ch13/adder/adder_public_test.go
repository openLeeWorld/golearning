package adder_test // 공용 api test용 패키지이름
// 노출된 함수, 메서드, 타입, 상수 및 변수를 통해서만 상호작용 강제
import (
	"test_examples/adder" // 패키지 import 필요
	"testing"
)

// Test<공용API함수> 이름으로 정해놓음
func TestAddNumbers(t *testing.T) {
	result := adder.AddNumbers(2, 3) // 공용 api 함수 불러옴
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
