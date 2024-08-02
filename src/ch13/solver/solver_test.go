package solver

import (
	"context"
	"errors"
	"strings"
	"testing"
)

// MathSolver를 위한 임시 객체 타입
type MathSolverStub struct{}

// 스텁이 활용하는 Resolve 메서드
func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
	switch expr {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10":
		return 0, errors.New("invalid expression: ( 2 + 2 * 10")
	}
	return 0, nil
}

func TestProcessor_ProcessExpressions(t *testing.T) {
	p := Processor{MathSolverStub{}}
	in := strings.NewReader(`2 + 2 * 10
( 2 + 2 ) * 10
( 2 + 2 * 10`) // 각 문장을 \n으로 구분함
	data := []float64{22, 40, 0, 0}
	for _, d := range data {
		result, err := p.ProcessExpression(context.Background(), in)
		if err != nil {
			t.Error(err)
		}
		if result != d {
			t.Errorf("Expected result %f, got %f", d, result)
		}
	}
}
