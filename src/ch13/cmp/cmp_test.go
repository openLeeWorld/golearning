package cmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  37,
	}
	result := CreatePerson("Dennis", 37)

	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	}) // 시간값 차이로 인한 자체 비교 함수 정의

	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	} // 일치하지 않으면 일치하지 않는 부분을 기술하는 문자열을 반환한다.
}
