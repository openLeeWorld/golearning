package main

import "fmt"

// Add 함수는 두 개의 인자를 받아서 더한 값을 반환합니다.
// T는 숫자 타입이어야 합니다.
func Add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	// 정수형으로 Add 함수 호출
	intResult := Add(3, 4)
	fmt.Printf("intResult: %d\n", intResult)

	// 실수형으로 Add 함수 호출
	floatResult := Add(3.5, 4.2)
	fmt.Printf("floatResult: %.2f\n", floatResult)
}
