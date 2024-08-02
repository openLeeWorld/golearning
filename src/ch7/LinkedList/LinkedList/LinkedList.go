package LinkedList

type LinkedList struct {
	Value interface{} // 사용자 정의된 제네릭 대신 빈 인터페이스로 값 받음
	Next  *LinkedList
}

// 슬라이스, 배열, 맵 이외의 데이터 구조가 필요하고 단일 타입에서만 동작x하려면
func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		} // 헤드 처리
	}

	if ll.Next == nil {
		ll.Next = &LinkedList{
			Value: val,
		}
	} else {
		ll.Next = ll.Next.Insert(pos-1, val)
	}

	return ll
}
