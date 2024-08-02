package adder

import (
	"fmt"
	"testing"
)

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		//t.Error("incorrect result: expected 5, got", result)
		t.Errorf("incorrect result: expected %d, got %s", 5, fmt.Sprint(result))
	} // 실패가 발견되는대로 멈춰야 하면 Fatal과 Fatalf 사용
}
