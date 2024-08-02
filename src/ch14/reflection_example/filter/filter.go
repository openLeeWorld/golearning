package filter

import "reflect"

// 다음은 리플렉션을 활용해 모든 타입에 대해 필터링을 수행하는 함수이다.
func Filter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			out = reflect.Append(out, curVal)
		} // 필터링 된 결과에 따라 결과에 현재 값을 넣는다.
	}
	return out.Interface()
}

func FilterString(s []string, f func(string) bool) []string {
	out := make([]string, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func FilterInt(s []int, f func(int) bool) []int {
	out := make([]int, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}
