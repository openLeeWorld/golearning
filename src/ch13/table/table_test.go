package table

import "testing"

func TestDoMathTable(t *testing.T) {
	data := []struct { // 테스트 테이블 항목과 데이터 선언
		name     string // 각 테스트 이름
		num1     int    // 인수1
		num2     int    // 인수2
		op       string // 인수3
		expected int    // 테스트별 실제 기대값
		errMsg   string // 테스트별 에러메시지
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"subtraction", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_division", 2, 0, "/", 0, `division by zero`},
		{"bad_op", 2, 2, "?", 0, `unknown operator ?`},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) { // 반복되는 패턴을 테이블로
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
