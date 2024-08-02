package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	child, cancel2 := context.WithTimeout(parent, 3*time.Second)
	defer cancel2()
	start := time.Now()
	<-child.Done() // context의 Done 메서드
	end := time.Now()
	fmt.Println(end.Sub(start)) // 자식 컨텍스트는 부모 컨텍스트의 시간 제한에 걸린다.
} // 답: 2초
